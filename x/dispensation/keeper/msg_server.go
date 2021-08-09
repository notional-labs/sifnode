package keeper

import (
	"context"
	"fmt"
	"strconv"

	dispensationUtils "github.com/Sifchain/sifnode/x/dispensation/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	"github.com/Sifchain/sifnode/x/dispensation/types"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the clp MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (srv msgServer) CreateDistribution(ctx context.Context,
	msg *types.MsgCreateDistribution) (*types.MsgCreateDistributionResponse, error) {

	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("sifnode"),
		tracer.WithEnv("test"),
	)
	defer tracer.Stop()

	// // Start a root span.
	span := tracer.StartSpan("disp.createserver")
	defer span.Finish()

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	distributionName := fmt.Sprintf("%d_%s", sdkCtx.BlockHeight(), msg.Distributor)
	verifySpan := tracer.StartSpan("disp.VerifyAndSetDistribution", tracer.ChildOf(span.Context()))
	// Verify if distribution already exists
	err := srv.Keeper.VerifyAndSetDistribution(sdkCtx, distributionName, msg.DistributionType, msg.AuthorizedRunner)
	verifySpan.Finish(tracer.WithError(err))
	if err != nil {
		return nil, err
	}

	totalOutput, err := dispensationUtils.TotalOutput(msg.Output)
	if err != nil {
		return nil, errors.Wrap(err, "Error calculating required amount from outputs")
	}
	accumSpan := tracer.StartSpan("disp.AccumulateDrops", tracer.ChildOf(span.Context()))
	//Accumulate all Drops into the ModuleAccount
	err = srv.Keeper.AccumulateDrops(sdkCtx, msg.Distributor, totalOutput)
	accumSpan.Finish(tracer.WithError(err))
	if err != nil {
		return nil, err
	}
	createSpan := tracer.StartSpan("disp.CreateDrops", tracer.ChildOf(span.Context()))
	//Create drops and Store Historical Data
	err = srv.Keeper.CreateDrops(sdkCtx, msg.Output, distributionName, msg.DistributionType, msg.AuthorizedRunner)
	createSpan.Finish(tracer.WithError(err))
	if err != nil {
		return nil, err
	}

	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDistributionStarted,
			sdk.NewAttribute(types.AttributeKeyFromModuleAccount, types.GetDistributionModuleAddress().String()),
			sdk.NewAttribute(types.AttributeKeyDistributionName, distributionName),
			sdk.NewAttribute(types.AttributeKeyDistributionType, msg.DistributionType.String()),
		),
	})

	return &types.MsgCreateDistributionResponse{}, nil
}

func (srv msgServer) CreateUserClaim(ctx context.Context,
	claim *types.MsgCreateUserClaim) (*types.MsgCreateClaimResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if srv.Keeper.ExistsClaim(sdkCtx, claim.UserClaimAddress, claim.UserClaimType) {
		sdkCtx.Logger().Info("Claim already exists for user :", claim.UserClaimAddress)
		return nil, errors.Wrap(types.ErrInvalid, "Claim already exists for user")
	}
	newClaim, err := types.NewUserClaim(claim.UserClaimAddress, claim.UserClaimType, sdkCtx.BlockTime())
	if err != nil {
		return nil, err
	}
	err = srv.Keeper.SetClaim(sdkCtx, newClaim)
	if err != nil {
		return nil, err
	}
	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeClaimCreated,
			sdk.NewAttribute(types.AttributeKeyClaimUser, newClaim.UserAddress),
			sdk.NewAttribute(types.AttributeKeyClaimType, newClaim.UserClaimType.String()),
			sdk.NewAttribute(types.AttributeKeyClaimTime, newClaim.UserClaimTime.String()),
		),
	})
	return &types.MsgCreateClaimResponse{}, nil
}

func (srv msgServer) RunDistribution(ctx context.Context, distribution *types.MsgRunDistribution) (*types.MsgRunDistributionResponse, error) {
	// Not checking whether the distribution exists or not .
	// We only need to run and execute distribution records
	// Distribute 10 drops for msg.DistributionName authorized to msg.DistributionRunner

	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("sifnode"),
		tracer.WithEnv("test"),
	)
	defer tracer.Stop()

	// // Start a root span.
	span := tracer.StartSpan("disp.createserver")
	defer span.Finish()

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	distributeSpan := tracer.StartSpan("disp.DistributeDrops", tracer.ChildOf(span.Context()))
	records, err := srv.Keeper.DistributeDrops(sdkCtx, sdkCtx.BlockHeight(), distribution.DistributionName, distribution.AuthorizedRunner, distribution.DistributionType)
	distributeSpan.Finish(tracer.WithError(err))
	if err != nil {
		return nil, err
	}

	recordEvents := make([]sdk.Event, len(records.DistributionRecords)+1)
	for i, record := range records.DistributionRecords {
		ev := sdk.NewEvent(
			types.EventTypeDistributionRecordsList+strconv.Itoa(i),
			sdk.NewAttribute(types.AttributeKeyDistributionRecordAddress, record.RecipientAddress),
			sdk.NewAttribute(types.AttributeKeyDistributionRecordType, record.DistributionType.String()),
			sdk.NewAttribute(types.AttributeKeyDistributionRecordAmount, record.Coins.String()),
		)
		recordEvents[i] = ev
	}
	recordEvents[len(recordEvents)-1] = sdk.NewEvent(
		types.EventTypeDistributionRun,
		sdk.NewAttribute(types.AttributeKeyDistributionName, distribution.DistributionName),
		sdk.NewAttribute(types.AttributeKeyDistributionRunner, distribution.AuthorizedRunner),
	)

	sdkCtx.EventManager().EmitEvents(recordEvents)
	return &types.MsgRunDistributionResponse{}, nil
}
