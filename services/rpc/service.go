package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/lengwh/dapp-wallet-sign/hsm"
	"github.com/lengwh/dapp-wallet-sign/leveldb"
	"github.com/lengwh/dapp-wallet-sign/protobuf/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync/atomic"
)

const MaxReceivedMessageSize = 1024 * 1024 * 30 // 30 MB

type RpcServerConfig struct {
	GrpcHost  string
	GrpcPort  string
	KeyPath   string
	keyName   string
	HsmEnable bool
}

type RpcServer struct {
	*RpcServerConfig
	db        *leveldb.Keys
	HsmClient *hsm.HsmClient // Assuming HsmClient is defined elsewhere
	wallet.UnimplementedWalletServiceServer
	stopped atomic.Bool
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func NewRpcServer(db *leveldb.Keys, config *RpcServerConfig) (*RpcServer, error) {
	hsmClient, err := hsm.NewHsmClient(context.Background(), config.KeyPath, config.keyName)
	if err != nil {
		log.Error("new hsm client fail", "err", err)
		return nil, err
	}
	return &RpcServer{
		RpcServerConfig: config,
		db:              db,
		HsmClient:       hsmClient,
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {

	go func(s *RpcServer) {
		listener, err := net.Listen("tcp", s.GrpcHost+":"+s.GrpcPort)
		if err != nil {
			log.Error("failed to listen", "error", err)
		}
		opt := grpc.MaxRecvMsgSize(MaxReceivedMessageSize)

		gs := grpc.NewServer(opt, grpc.ChainStreamInterceptor(nil))

		reflection.Register(gs)
		wallet.RegisterWalletServiceServer(gs, s)
		log.Info("starting rpc services", "address", s.GrpcHost+":"+s.GrpcPort)
		if err := gs.Serve(listener); err != nil {
			log.Error("failed to serve", "error", err)
		}
	}(s)
	return nil
}
