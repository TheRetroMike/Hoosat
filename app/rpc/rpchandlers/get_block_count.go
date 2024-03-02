package rpchandlers

import (
	"github.com/Hoosat-Oy/hoosatd/app/appmessage"
	"github.com/Hoosat-Oy/hoosatd/app/rpc/rpccontext"
	"github.com/Hoosat-Oy/hoosatd/infrastructure/network/netadapter/router"
)

// HandleGetBlockCount handles the respectively named RPC command
func HandleGetBlockCount(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	syncInfo, err := context.Domain.Consensus().GetSyncInfo()
	if err != nil {
		return nil, err
	}
	response := appmessage.NewGetBlockCountResponseMessage(syncInfo)
	return response, nil
}
