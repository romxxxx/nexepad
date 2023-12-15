package client

import (
	"context"
	"time"

	"github.com/romxxxx/nexepad/cmd/nexepawallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/romxxxx/nexepad/cmd/nexepawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the nexepawalletd server, and returns the client instance
func Connect(address string) (pb.nexepawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("nexepawallet daemon is not running, start it with `nexepawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewnexepawalletdClient(conn), func() {
		conn.Close()
	}, nil
}
