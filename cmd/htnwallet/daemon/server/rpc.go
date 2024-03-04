package server

import (
	"time"

	"github.com/Hoosat-Oy/htnd/domain/dagconfig"
	"github.com/Hoosat-Oy/htnd/infrastructure/network/rpcclient"
)

func connectToRPC(params *dagconfig.Params, rpcServer string, timeout uint32) (*rpcclient.RPCClient, error) {
	rpcAddress, err := params.NormalizeRPCServerAddress(rpcServer)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpcclient.NewRPCClient(rpcAddress)
	if err != nil {
		return nil, err
	}

	if timeout != 0 {
		rpcClient.SetTimeout(time.Duration(timeout) * time.Second)
	}

	return rpcClient, err
}
