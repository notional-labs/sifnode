package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func SetupHandlers(app *SifchainApp) {
	app.UpgradeKeeper.SetUpgradeHandler("0.9.0", func(ctx sdk.Context, plan types.Plan) {})
	app.UpgradeKeeper.SetUpgradeHandler("0.9.1-rc.10", func(ctx sdk.Context, plan types.Plan) {})
}
