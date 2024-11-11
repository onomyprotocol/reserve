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
mnemonic1="top toddler wrist parade hobby supply odor ginger resource copy square tell vanish pride volcano effort planet style transfer pipe wise bus tuition luxury"
mnemonic2="panther giant oyster hand song region chunk coil laundry glance ball denial void ramp palm fiscal pizza soccer before upset diet valid story cement"
mnemonic3="gap track crop knee galaxy square case resemble subway math moon mom casino trade finish exotic author comic gap margin elegant claw fire business"

echo $mnemonic1| reserved keys add validator1 --recover --keyring-backend=test --home=$HOME/.reserved/validator1
echo $mnemonic2| reserved keys add validator2 --recover --keyring-backend=test --home=$HOME/.reserved/validator2
echo $mnemonic3| reserved keys add validator3 --recover --keyring-backend=test --home=$HOME/.reserved/validator3

# create validator node with tokens to transfer to the three other nodes
reserved genesis add-genesis-account $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator1 
reserved genesis add-genesis-account $(reserved keys show validator2 -a --keyring-backend=test --home=$HOME/.reserved/validator2) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator1 
reserved genesis add-genesis-account $(reserved keys show validator3 -a --keyring-backend=test --home=$HOME/.reserved/validator3) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator1 
reserved genesis add-genesis-account $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator2 
reserved genesis add-genesis-account $(reserved keys show validator2 -a --keyring-backend=test --home=$HOME/.reserved/validator2) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator2 
reserved genesis add-genesis-account $(reserved keys show validator3 -a --keyring-backend=test --home=$HOME/.reserved/validator3) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator2 
reserved genesis add-genesis-account $(reserved keys show validator1 -a --keyring-backend=test --home=$HOME/.reserved/validator1) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator3 
reserved genesis add-genesis-account $(reserved keys show validator2 -a --keyring-backend=test --home=$HOME/.reserved/validator2) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator3 
reserved genesis add-genesis-account $(reserved keys show validator3 -a --keyring-backend=test --home=$HOME/.reserved/validator3) 10000000000000000000000000000000stake,10000000000000000000000000000000usdt,10000000000000000000000000000000atom --home=$HOME/.reserved/validator3 
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
reserved q gov proposals
# reserved tx gov submit-legacy-proposal active-collateral "title" "description" "atom" "10" "0.1" "10000" 10000000000000000000stake --keyring-backend=test  --home=$HOME/.reserved/validator1 --from validator1 -y --chain-id testing-1 --fees 20stake

reserved tx gov submit-proposal ./script/proposal-vault-1.json --home=$HOME/.reserved/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

# # vote
sleep 7
reserved tx gov vote 1 yes  --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake
reserved tx gov vote 1 yes  --from validator2 --keyring-backend test --home ~/.reserved/validator2 --chain-id testing-1 -y --fees 20stake
reserved tx gov vote 1 yes  --from validator3 --keyring-backend test --home ~/.reserved/validator3 --chain-id testing-1 -y --fees 20stake

# wait voting_perio=15s
echo "========sleep=========="
sleep 15
reserved q gov proposals
reserved tx oracle set-price usdt 1  --home=$HOME/.reserved/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y
reserved tx oracle set-price atom 8  --home=$HOME/.reserved/validator2  --from validator2 --keyring-backend test --fees 20stake --chain-id testing-1 -y

sleep 7

reserved tx vaults create-vault 1250000000usdt 50000000nomUSD --home=$HOME/.reserved/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y
reserved tx vaults create-vault 1250000000atom 50000000nomUSD --home=$HOME/.reserved/validator2  --from validator2 --keyring-backend test --fees 20stake --chain-id testing-1 -y
# killall reserved || true


# 02420E85EC300BE3E9219C5D8330207451F0D33764
# liquidationMap map[]
# nextId 0 <nil>
# 10:26AM ERR panic recovered in runTx err="recovered: module account reserve does not exist: unknown address [cosmos/cosmos-sdk@v0.50.6/x/bank/keeper/keeper.go:286]\nstack:\ngoroutine 438 [running]:\nruntime/debug.Stack()\n\t/usr/local/go/src/runtime/debug/stack.go:26 +0x64\ngithub.com/cosmos/cosmos-sdk/baseapp.NewBaseApp.newDefaultRecoveryMiddleware.func5({0x106011480, 0x140023cc840})\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/recovery.go:74 +0x24\ngithub.com/cosmos/cosmos-sdk/baseapp.NewBaseApp.newDefaultRecoveryMiddleware.newRecoveryMiddleware.func7({0x106011480?, 0x140023cc840?})\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/recovery.go:42 +0x38\ngithub.com/cosmos/cosmos-sdk/baseapp.processRecovery({0x106011480, 0x140023cc840}, 0x14003c493d8?)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/recovery.go:31 +0x38\ngithub.com/cosmos/cosmos-sdk/baseapp.processRecovery({0x106011480, 0x140023cc840}, 0x8?)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/recovery.go:36 +0x60\ngithub.com/cosmos/cosmos-sdk/baseapp.(*BaseApp).runTx.func1()\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/baseapp.go:837 +0xd4\npanic({0x106011480?, 0x140023cc840?})\n\t/usr/local/go/src/runtime/panic.go:785 +0x124\ngithub.com/cosmos/cosmos-sdk/x/bank/keeper.BaseKeeper.SendCoinsFromModuleToModule({{{{0x106376aa0, 0x14000d05160}, {0x1063259e0, 0x14001588318}, {0x10637e670, 0x14000355400}, {0x106364930, 0x1400208e620}, {0x140015946a8, {...}, ...}, ...}, ...}, ...}, ...)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/x/bank/keeper/keeper.go:286 +0x15c\ngithub.com/onomyprotocol/reserve/x/vaults/keeper.(*Keeper).CreateNewVault(0x14003c4aab8, {0x10635c228, 0x14006ab5080}, {0x140023664b0, 0x14, 0x15}, {{0x140019fcb80?, 0x1063259c0?}, {0x14001b59520?}}, {{0x140019fcbc0?, ...}, ...})\n\t/Users/donglieu/102024/main/reserve/x/vaults/keeper/vault.go:68 +0x3dc\ngithub.com/onomyprotocol/reserve/x/vaults/keeper.msgServer.CreateVault({{{0x106376aa0, 0x14000d05160}, {0x1063259e0, 0x140015883a0}, {0x1519327e8, 0x14000595688}, {0x10633ca60, 0x140015fd040}, {0x106336280, 0x14000be1c78}, ...}}, ...)\n\t/Users/donglieu/102024/main/reserve/x/vaults/keeper/msg_server.go:67 +0x74\ngithub.com/onomyprotocol/reserve/x/vaults/types._Msg_CreateVault_Handler.func1({0x10635c228?, 0x14006ab5080?}, {0x1061a5b00?, 0x14006296080?})\n\t/Users/donglieu/102024/main/reserve/x/vaults/types/tx.pb.go:1102 +0xd0\ngithub.com/cosmos/cosmos-sdk/baseapp.(*MsgServiceRouter).registerMsgServiceHandler.func2.1({0x10635c0d8, 0x14003c68008}, {0x14003c4b098?, 0x103f592f0?}, 0x358?, 0x140011a5bc0)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/msg_service_router.go:175 +0x98\ngithub.com/onomyprotocol/reserve/x/vaults/types._Msg_CreateVault_Handler({0x1062c7640, 0x140012ca008}, {0x10635c0d8, 0x14003c68008}, 0x1063132e8, 0x14001b59580)\n\t/Users/donglieu/102024/main/reserve/x/vaults/types/tx.pb.go:1104 +0x148\ngithub.com/cosmos/cosmos-sdk/baseapp.(*MsgServiceRouter).registerMsgServiceHandler.func2({{0x10635c0a0, 0x107e7b100}, {0x106377360, 0x14006296100}, {{0x0, 0x0}, {0x14001c322a0, 0x9}, 0x67, {0x60f9708, ...}, ...}, ...}, ...)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/msg_service_router.go:198 +0x2b0\ngithub.com/cosmos/cosmos-sdk/baseapp.(*BaseApp).runMsgs(_, {{0x10635c0a0, 0x107e7b100}, {0x106377360, 0x14006296100}, {{0x0, 0x0}, {0x14001c322a0, 0x9}, 0x67, ...}, ...}, ...)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/baseapp.go:1010 +0x170\ngithub.com/cosmos/cosmos-sdk/baseapp.(*BaseApp).runTx(0x140023fc248, 0x7, {0x140014f4140, 0x123, 0x123})\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/baseapp.go:948 +0xbf8\ngithub.com/cosmos/cosmos-sdk/baseapp.(*BaseApp).deliverTx(0x140023fc248, {0x140014f4140?, 0x123?, 0x14001a01540?})\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/baseapp.go:763 +0x88\ngithub.com/cosmos/cosmos-sdk/baseapp.(*BaseApp).internalFinalizeBlock(0x140023fc248, {0x10635c0a0, 0x107e7b100}, 0x14003c33b00)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/abci.go:790 +0xc2c\ngithub.com/cosmos/cosmos-sdk/baseapp.(*BaseApp).FinalizeBlock(0x140023fc248, 0x14003c33b00)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/baseapp/abci.go:884 +0x118\ngithub.com/cosmos/cosmos-sdk/server.cometABCIWrapper.FinalizeBlock(...)\n\t/Users/donglieu/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/server/cmt_abci.go:44\ngithub.com/cometbft/cometbft/abci/client.(*localClient).FinalizeBlock(0x14001688c20?, {0x10635c308?, 0x107e7b100?}, 0x10a8b0a68?)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/abci/client/local_client.go:185 +0xe4\ngithub.com/cometbft/cometbft/proxy.(*appConnConsensus).FinalizeBlock(0x14001a040c0, {0x10635c308, 0x107e7b100}, 0x14003c33b00)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/proxy/app_conn.go:104 +0x124\ngithub.com/cometbft/cometbft/state.(*BlockExecutor).applyBlock(_, {{{0xb, 0x0}, {0x14000a14069, 0x7}}, {0x14000a14080, 0x9}, 0x1, 0x66, {{0x140061a1b20, ...}, ...}, ...}, ...)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/state/execution.go:224 +0x410\ngithub.com/cometbft/cometbft/state.(*BlockExecutor).ApplyVerifiedBlock(...)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/state/execution.go:202\ngithub.com/cometbft/cometbft/consensus.(*State).finalizeCommit(0x1400248b508, 0x67)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:1772 +0x97c\ngithub.com/cometbft/cometbft/consensus.(*State).tryFinalizeCommit(0x1400248b508, 0x67)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:1682 +0x26c\ngithub.com/cometbft/cometbft/consensus.(*State).enterCommit.func1()\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:1617 +0x8c\ngithub.com/cometbft/cometbft/consensus.(*State).enterCommit(0x1400248b508, 0x67, 0x0)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:1655 +0xac0\ngithub.com/cometbft/cometbft/consensus.(*State).addVote(0x1400248b508, 0x14001334750, {0x14001004090, 0x28})\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:2335 +0x182c\ngithub.com/cometbft/cometbft/consensus.(*State).tryAddVote(0x1400248b508, 0x14001334750, {0x14001004090?, 0x103dd81c4?})\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:2067 +0x28\ngithub.com/cometbft/cometbft/consensus.(*State).handleMsg(0x1400248b508, {{0x10632e320, 0x14000e82200}, {0x14001004090, 0x28}})\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:929 +0x2fc\ngithub.com/cometbft/cometbft/consensus.(*State).receiveRoutine(0x1400248b508, 0x0)\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:836 +0x2b8\ncreated by github.com/cometbft/cometbft/consensus.(*State).OnStart in goroutine 372\n\t/Users/donglieu/go/pkg/mod/github.com/cometbft/cometbft@v0.38.12/consensus/state.go:398 +0xf0\n: panic" module=server