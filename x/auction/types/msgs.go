package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	Query_serviceDesc = _Query_serviceDesc
	Msg_serviceDesc   = _Msg_serviceDesc
)

func NewMsgBid(addr string, auctionID uint64, amount sdk.Coin, ReciveRate string) MsgBid {
	return MsgBid{
		Bidder:     addr,
		AuctionId:  auctionID,
		ReciveRate: ReciveRate,
		Amount:     amount,
	}
}

func NewMsgCancelBid(bider string, bidID, auctionID uint64) MsgCancelBid {
	return MsgCancelBid{
		Bidder:    bider,
		BidId:     bidID,
		AuctionId: auctionID,
	}
}
