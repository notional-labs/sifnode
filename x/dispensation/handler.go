package dispensation

import (
	"fmt"

	"github.com/Sifchain/sifnode/x/dispensation/keeper"
	"github.com/Sifchain/sifnode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// NewHandler creates an sdk.Handler for all the dispensation type messages
func NewHandler(k Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		rules := []tracer.SamplingRule{tracer.RateRule(1)}
		tracer.Start(
			tracer.WithSamplingRules(rules),
			tracer.WithService("sifnode"),
			tracer.WithEnv("test"),
		)
		defer tracer.Stop()

		// // Start a root span.
		span := tracer.StartSpan("disp.handler")
		defer span.Finish()

		switch msg := msg.(type) {
		case *types.MsgCreateDistribution:
			// // Create a child of it, computing the time needed to read a file.
			createSpan := tracer.StartSpan("disp.create", tracer.ChildOf(span.Context()))
			res, err := msgServer.CreateDistribution(sdk.WrapSDKContext(ctx), msg)
			createSpan.Finish(tracer.WithError(err))
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateUserClaim:
			res, err := msgServer.CreateUserClaim(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRunDistribution:
			runSpan := tracer.StartSpan("disp.run", tracer.ChildOf(span.Context()))
			res, err := msgServer.RunDistribution(sdk.WrapSDKContext(ctx), msg)
			runSpan.Finish(tracer.WithError(err))
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
