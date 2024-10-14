set -xeu
MNE1="amused rural desk trick safe whip first menu worth swap enhance punch spin figure elevator abandon camera idea peace nurse coyote adjust modify produce"

# init
rm -rf .onomy
reserved init asd --chain-id local_onomy-1 --home .onomy -o

# provider and consumer keys
echo $MNE1 | reserved keys add god --recover --keyring-backend test --home .onomy

# provider validator
reserved genesis add-genesis-account god 1000000000000000000000000000anom --keyring-backend test  --home .onomy
reserved genesis gentx god 1000000000000000000000000anom  --keyring-backend test  --home .onomy --chain-id local_onomy-1
reserved genesis collect-gentxs --home .onomy

# provider genesis
cat .onomy/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="100"'                 > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="6s"'             > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["amount"]="100"' > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="anom"'                  > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="anom"'                  > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="anom"' > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="anom"'                > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json
cat .onomy/config/genesis.json | jq '.app_state["oracle"]["band_oracle_requests"]+=[{"request_id":1,"oracle_script_id":360,"symbols":["BTC","ETH","BAND","USDT"],"ask_count":4,"min_count":3,"prepare_gas":100000,"execute_gas":500000,"fee_limit":[{"denom":"uband","amount":"250000"}]}]'                > .onomy/config/tmp_genesis.json && mv .onomy/config/tmp_genesis.json .onomy/config/genesis.json