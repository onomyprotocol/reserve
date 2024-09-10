#!/bin/bash
set -xeu

# always returns true so set -e doesn't exit if it is not running.
killall reserved || true
rm -rf $HOME/.reserved/

mkdir $HOME/.reserved
mkdir $HOME/.reserved/validator1
mkdir $HOME/.reserved/validator2
mkdir $HOME/.reserved/validator3

# init all three validators
reserved init --chain-id=testing-1 validator1 --home=$HOME/.reserved/validator1
reserved init --chain-id=testing-1 validator2 --home=$HOME/.reserved/validator2
reserved init --chain-id=testing-1 validator3 --home=$HOME/.reserved/validator3

# create keys for all three validators
echo $(cat ./script/mnemonic/mnemonic1)| reserved keys add validator1 --recover --keyring-backend=test --home=$HOME/.reserved/validator1
echo $(cat ./script/mnemonic/mnemonic2)| reserved keys add validator2 --recover --keyring-backend=test --home=$HOME/.reserved/validator2
echo $(cat ./script/mnemonic/mnemonic3)| reserved keys add validator3 --recover --keyring-backend=test --home=$HOME/.reserved/validator3

# create validator node with tokens to transfer to the three other nodes
reserved genesis add-genesis-account $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator1 
reserved genesis add-genesis-account $(reserved keys show validator2 -a --keyring-backend=test --home=$HOME/.reserved/validator2) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator1 
reserved genesis add-genesis-account $(reserved keys show validator3 -a --keyring-backend=test --home=$HOME/.reserved/validator3) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator1 
reserved genesis add-genesis-account $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator2 
reserved genesis add-genesis-account $(reserved keys show validator2 -a --keyring-backend=test --home=$HOME/.reserved/validator2) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator2 
reserved genesis add-genesis-account $(reserved keys show validator3 -a --keyring-backend=test --home=$HOME/.reserved/validator3) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator2 
reserved genesis add-genesis-account $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator3 
reserved genesis add-genesis-account $(reserved keys show validator2 -a --keyring-backend=test --home=$HOME/.reserved/validator2) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator3 
reserved genesis add-genesis-account $(reserved keys show validator3 -a --keyring-backend=test --home=$HOME/.reserved/validator3) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt --home=$HOME/.reserved/validator3 
reserved genesis gentx validator1 1000000000000000000000stake --keyring-backend=test --home=$HOME/.reserved/validator1 --chain-id=testing-1
reserved genesis gentx validator2 1000000000000000000000stake --keyring-backend=test --home=$HOME/.reserved/validator2 --chain-id=testing-1
reserved genesis gentx validator3 1000000000000000000000stake --keyring-backend=test --home=$HOME/.reserved/validator3 --chain-id=testing-1

# cp validator2/config/gentx/*.json $HOME/.reserved/validator1/config/gentx/
# cp validator3/config/gentx/*.json $HOME/.reserved/validator1/config/gentx/
reserved genesis collect-gentxs --home=$HOME/.reserved/validator1 

# change app.toml values
VALIDATOR1_APP_TOML=$HOME/.reserved/validator1/config/app.toml
VALIDATOR2_APP_TOML=$HOME/.reserved/validator2/config/app.toml
VALIDATOR3_APP_TOML=$HOME/.reserved/validator3/config/app.toml

# validator1
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9050|g' $VALIDATOR1_APP_TOML
sed -i -E 's|127.0.0.1:9090|127.0.0.1:9050|g' $VALIDATOR1_APP_TOML
sed -i -E 's|minimum-gas-prices = ""|minimum-gas-prices = "0.0001stake"|g' $VALIDATOR1_APP_TOML

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $VALIDATOR2_APP_TOML
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $VALIDATOR2_APP_TOML
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $VALIDATOR2_APP_TOML
sed -i -E 's|minimum-gas-prices = ""|minimum-gas-prices = "0.0001stake"|g' $VALIDATOR2_APP_TOML

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $VALIDATOR3_APP_TOML
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $VALIDATOR3_APP_TOML
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $VALIDATOR3_APP_TOML
sed -i -E 's|minimum-gas-prices = ""|minimum-gas-prices = "0.0001stake"|g' $VALIDATOR3_APP_TOML

# change config.toml values
VALIDATOR1_CONFIG=$HOME/.reserved/validator1/config/config.toml
VALIDATOR2_CONFIG=$HOME/.reserved/validator2/config/config.toml
VALIDATOR3_CONFIG=$HOME/.reserved/validator3/config/config.toml


# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $VALIDATOR1_CONFIG
sed -i -E 's|prometheus = false|prometheus = true|g' $VALIDATOR1_CONFIG


# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $VALIDATOR2_CONFIG
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26654|g' $VALIDATOR2_CONFIG
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26653|g' $VALIDATOR2_CONFIG
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $VALIDATOR2_CONFIG
sed -i -E 's|prometheus = false|prometheus = true|g' $VALIDATOR2_CONFIG
sed -i -E 's|prometheus_listen_addr = ":26660"|prometheus_listen_addr = ":26630"|g' $VALIDATOR2_CONFIG

# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $VALIDATOR3_CONFIG
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26651|g' $VALIDATOR3_CONFIG
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $VALIDATOR3_CONFIG
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $VALIDATOR3_CONFIG
sed -i -E 's|prometheus = false|prometheus = true|g' $VALIDATOR3_CONFIG
sed -i -E 's|prometheus_listen_addr = ":26660"|prometheus_listen_addr = ":26620"|g' $VALIDATOR3_CONFIG

# copy, update validator1 genesis file to validator2-3
update_test_genesis () {
    cat $HOME/.reserved/validator1/config/genesis.json | jq "$1" > tmp.json && mv tmp.json $HOME/.reserved/validator1/config/genesis.json
}

update_test_genesis '.app_state["gov"]["params"]["voting_period"] = "15s"'
update_test_genesis '.app_state["gov"]["params"]["expedited_voting_period"] = "10s"'

cp $HOME/.reserved/validator1/config/genesis.json $HOME/.reserved/validator2/config/genesis.json
cp $HOME/.reserved/validator1/config/genesis.json $HOME/.reserved/validator3/config/genesis.json

# copy tendermint node id of validator1 to persistent peers of validator2-3
node1=$(reserved tendermint show-node-id --home=$HOME/.reserved/validator1)
node2=$(reserved tendermint show-node-id --home=$HOME/.reserved/validator2)
node3=$(reserved tendermint show-node-id --home=$HOME/.reserved/validator3)
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$node1@localhost:26656,$node2@localhost:26656,$node3@localhost:26656\"|g" $HOME/.reserved/validator1/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$node1@localhost:26656,$node2@localhost:26656,$node3@localhost:26656\"|g" $HOME/.reserved/validator2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$node1@localhost:26656,$node2@localhost:26656,$node3@localhost:26656\"|g" $HOME/.reserved/validator3/config/config.toml


# # start all three validators
screen -S onomy1 -t onomy1 -d -m reserved start --home=$HOME/.reserved/validator1
screen -S onomy2 -t onomy2 -d -m reserved start --home=$HOME/.reserved/validator2
screen -S onomy3 -t onomy3 -d -m reserved start --home=$HOME/.reserved/validator3

# submit proposal add usdt
sleep 7
reserved tx gov submit-legacy-proposal add-stable-coin "d" "d" "usdt" "100000000000000000000000000000" "1" "0.001" "0.001" $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1)  10000000000000000000stake --keyring-backend=test  --home=$HOME/.reserved/validator1 --from validator1 -y --chain-id testing-1 --fees 20stake

# vote
sleep 7
reserved tx gov vote 1 yes  --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake
reserved tx gov vote 1 yes  --from validator2 --keyring-backend test --home ~/.reserved/validator2 --chain-id testing-1 -y --fees 20stake
reserved tx gov vote 1 yes  --from validator3 --keyring-backend test --home ~/.reserved/validator3 --chain-id testing-1 -y --fees 20stake

# wait voting_perio=15s
sleep 15
echo "========sleep=========="

# check add usdt, balances
reserved q psm stablecoin usdt
reserved q bank balances onomy1wa3u4knw74r598quvzydvca42qsmk6jrc6uj7m

# tx swap usdt to ist
echo "========swap==========="
reserved tx psm swap-to-ist 100000000000000000000000000000usdt --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake

sleep 7
# Check account after swap
reserved q bank balances onomy1wa3u4knw74r598quvzydvca42qsmk6jrc6uj7m

# tx swap ist to usdt
reserved tx psm swap-to-stablecoin usdt 1000IST --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake

sleep 7
# Check account after swap
reserved q bank balances onomy1wa3u4knw74r598quvzydvca42qsmk6jrc6uj7m