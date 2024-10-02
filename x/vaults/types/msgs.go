package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	Query_serviceDesc = _Query_serviceDesc
	Msg_serviceDesc   = _Msg_serviceDesc
)

func NewMsgCreateVault(owner string, collateral, minted sdk.Coin) MsgCreateVault {
	return MsgCreateVault{
		Owner:      owner,
		Collateral: collateral,
		Minted:     minted,
	}
}

func NewMsgDeposit(vaultId uint64, sender string, amount sdk.Coin) MsgDeposit {
	return MsgDeposit{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgWithdraw(vaultId uint64, sender string, amount sdk.Coin) MsgWithdraw {
	return MsgWithdraw{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgMint(vaultId uint64, sender string, amount sdk.Coin) MsgMint {
	return MsgMint{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgRepay(vaultId uint64, sender string, amount sdk.Coin) MsgRepay {
	return MsgRepay{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgClose(vaultId uint64, sender string) MsgClose {
	return MsgClose{
		VaultId: vaultId,
		Sender:  sender,
	}
}
