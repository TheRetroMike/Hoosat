package client

import (
	"context"
	"time"

	"github.com/Hoosat-Oy/HTND/cmd/htnwallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/Hoosat-Oy/HTND/cmd/htnwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the hoosatwalletd server, and returns the client instance
func Connect(address string) (pb.HoosatwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("hoosatwallet daemon is not running, start it with `hoosatwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewHoosatwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
