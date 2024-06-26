package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/elys-network/elys/x/amm/types"
	assetprofiletypes "github.com/elys-network/elys/x/assetprofile/types"
	paramtypes "github.com/elys-network/elys/x/parameter/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	denom := req.Denom
	addr := req.Address
	address, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	balance := k.bankKeeper.GetBalance(ctx, address, denom)
	if denom != paramtypes.Elys {
		commitment := k.commitmentKeeper.GetCommitments(ctx, addr)
		claimed := commitment.GetClaimedForDenom(denom)
		commitBalance := sdk.NewCoin(denom, claimed)

		baseCurrency, found := k.assetProfileKeeper.GetUsdcDenom(ctx)
		if !found {
			return nil, errorsmod.Wrapf(assetprofiletypes.ErrAssetProfileNotFound, "asset %s not found", paramtypes.BaseCurrency)
		}

		// If it is USDC, we should add bank module balance as well.
		if denom == baseCurrency {
			balance = balance.Add(commitBalance)
		} else {
			balance = commitBalance
		}
	}

	return &types.QueryBalanceResponse{
		Balance: balance,
	}, nil
}
