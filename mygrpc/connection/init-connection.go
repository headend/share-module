package connection

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type RpcClient struct {
	Client *grpc.ClientConn
	Err error
}

func (rpc_cli *RpcClient) InitializeClient(host string, port uint16) *RpcClient {
	rpcServerAddr := fmt.Sprintf("%s:%d", host, port)
	rpc_cli.Client, rpc_cli.Err = grpc.Dial(rpcServerAddr,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(20 * 1024 * 1024)),
		grpc.WithInsecure())
	if rpc_cli.Err != nil {
		log.Fatalf("failed to connect server: %v", rpc_cli.Err)
	}
	return rpc_cli
}
