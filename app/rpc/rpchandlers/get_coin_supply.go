package rpchandlers

import (
	"github.com/romxxxx/nexepad/app/appmessage"
	"github.com/romxxxx/nexepad/app/rpc/rpccontext"
	"github.com/romxxxx/nexepad/domain/consensus/utils/constants"
	"github.com/romxxxx/nexepad/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when nexepad is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSompiSupply, err := context.UTXOIndex.GetCirculatingSompiSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSompi,
		circulatingSompiSupply,
	)

	return response, nil
}
