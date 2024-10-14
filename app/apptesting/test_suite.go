package apptesting

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/reserve/app"
)

type KeeperTestHelper struct {
	suite.Suite

	App      *app.App
	Ctx      sdk.Context
	TestAccs []sdk.AccAddress
}

var (
	baseTestAccts        = []sdk.AccAddress{}
	defaultTestStartTime = time.Now().UTC()
)

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

func init() {
	baseTestAccts = CreateRandomAccounts(3)
}

func (s *KeeperTestHelper) Setup() {
	s.App = app.Setup(s.T(), false)
	s.Ctx = s.App.BaseApp.NewContextLegacy(false, cmtproto.Header{Height: 1, ChainID: s.App.ChainID(), Time: defaultTestStartTime})

	vals, err := s.App.StakingKeeper.GetAllValidators(s.Ctx)

	s.TestAccs = []sdk.AccAddress{}
	s.TestAccs = append(s.TestAccs, baseTestAccts...)
	if err != nil {
		panic(err)
	}
	for _, val := range vals {
		var consAddr sdk.ConsAddress
		consAddr, _ = val.GetConsAddr()
		// newConsAddr := sdk.ConsAddress(pubkey.Address().Bytes())
		signingInfo := slashingtypes.NewValidatorSigningInfo(
			consAddr,
			s.Ctx.BlockHeight(),
			0,
			time.Unix(0, 0),
			false,
			0,
		)
		err := s.App.SlashingKeeper.SetValidatorSigningInfo(s.Ctx, consAddr, signingInfo)
		if err != nil {
			panic(err)
		}
	}
}

func (s *KeeperTestHelper) FundAccount(acccount sdk.AccAddress, moduleName string, coins sdk.Coins) {
	s.App.BankKeeper.MintCoins(s.Ctx, moduleName, coins)
	s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, moduleName, acccount, coins)
}
