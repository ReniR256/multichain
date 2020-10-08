package filecoin

import (
	"context"
	"fmt"
	"math/big"

	filaddress "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/renproject/pack"
)

// A GasEstimator returns the gas fee cap and gas premium that is needed in
// order to confirm transactions with an estimated maximum delay of one block.
// In distributed networks that collectively build, sign, and submit
// transactions, it is important that all nodes in the network have reached
// consensus on these values.
type GasEstimator struct {
	client   *Client
	gasLimit int64
}

// NewGasEstimator returns a simple gas estimator that fetches the ideal gas
// fee cap and gas premium for a filecoin transaction to be included in a block
// with minimal delay.
func NewGasEstimator(client *Client, gasLimit int64) GasEstimator {
	return GasEstimator{
		client:   client,
		gasLimit: gasLimit,
	}
}

// EstimateGasPrice returns gas fee cap and gas premium values being accepted
// in the filecoin chain at present. These numbers may vary as the chain's
// congestion level increases. It is safe to say that by using the fetched
// values, a transaction will be included in a block with minimal delay.
func (gasEstimator *GasEstimator) EstimateGasPrice(ctx context.Context) (pack.U256, pack.U256, error) {
	// Create a dummy "Send" message.
	msgIn := types.Message{
		Version:    types.MessageVersion,
		From:       filaddress.Undef,
		To:         filaddress.Undef,
		Value:      types.EmptyInt,
		Nonce:      0,
		GasLimit:   gasEstimator.gasLimit,
		GasFeeCap:  types.EmptyInt,
		GasPremium: types.EmptyInt,
		Method:     abi.MethodNum(0),
		Params:     []byte{},
	}

	// Estimate the gas fee cap and gas premium fields for this dummy message.
	msgOut, err := gasEstimator.client.node.GasEstimateMessageGas(ctx, &msgIn, &api.DefaultMessageSendSpec, types.EmptyTSK)
	if err != nil {
		return pack.NewU256([32]byte{}), pack.NewU256([32]byte{}), fmt.Errorf("estimating gas price: %v", err)
	}

	gasFeeCapBytes, err := msgOut.GasFeeCap.Bytes()
	if err != nil {
		return pack.NewU256([32]byte{}), pack.NewU256([32]byte{}), fmt.Errorf("getting abi/big bytes for %v: %v", msgOut.GasFeeCap, err)
	}
	gasPremiumBytes, err := msgOut.GasPremium.Bytes()
	if err != nil {
		return pack.NewU256([32]byte{}), pack.NewU256([32]byte{}), fmt.Errorf("getting abi/big bytes for %v: %v", msgOut.GasPremium, err)
	}
	gasFeeCap := big.NewInt(0).SetBytes(gasFeeCapBytes)
	gasPremium := big.NewInt(0).SetBytes(gasPremiumBytes)

	return pack.NewU256FromInt(gasFeeCap), pack.NewU256FromInt(gasPremium), nil
}
