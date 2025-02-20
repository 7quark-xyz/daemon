package prices

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/multicalls"
)

type TVaultToAsset struct {
	VaultAddress common.Address
	AssetAddress common.Address
	Value        *bigNumber.Int
}

/**************************************************************************************************
** fetchShareValueFromERC4626 will try to get the value of the assets for a ERC4626 vault type token
** It will return an array of struct with vault/asset/value
**************************************************************************************************/
func fetchShareValueFromERC4626(chainID uint64, blockNumber *uint64, tokenList []common.Address) []TVaultToAsset {
	vaultToAsset := []TVaultToAsset{}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokenList will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, tokenAddress := range tokenList {
		if token, ok := store.GetERC20(chainID, tokenAddress); ok {
			oneUnitScaledToDecimals := helpers.ToRawAmount(bigNumber.NewInt(1), token.Decimals)
			calls = append(calls, multicalls.GetConvertToAssets(tokenAddress.Hex(), tokenAddress, oneUnitScaledToDecimals))
			calls = append(calls, multicalls.GetAsset(tokenAddress.Hex(), tokenAddress))
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	var response map[string][]interface{}
	var blockNumberBigInt *big.Int

	if blockNumber == nil {
		currentBlockNumber, _ := ethereum.GetRPC(chainID).BlockNumber(context.Background())
		blockNumber = &currentBlockNumber
		response = multicalls.Perform(chainID, calls, nil)
	} else {
		blockNumberBigInt = big.NewInt(int64(*blockNumber))
		response = multicalls.Perform(chainID, calls, blockNumberBigInt)
	}

	for _, token := range tokenList {
		rawConvertedToAsset := response[token.Hex()+`convertToAssets`]
		rawAsset := response[token.Hex()+`asset`]
		if len(rawConvertedToAsset) == 0 || len(rawAsset) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawConvertedToAsset[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}

		vaultToAsset = append(vaultToAsset, TVaultToAsset{
			VaultAddress: token,
			AssetAddress: helpers.DecodeAddress(rawAsset),
			Value:        helpers.DecodeBigInt(rawConvertedToAsset),
		})
	}
	return vaultToAsset
}
