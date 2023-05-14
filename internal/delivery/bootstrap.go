package delivery

import (
	"context"
	"fmt"
	"github.com/morris-zheng/go-slim-core/discovery"
	"net"

	"github.com/morris-zheng/go-slim-micro-usersvc/export/user"
	userDelivery "github.com/morris-zheng/go-slim-micro-usersvc/internal/delivery/user"
	"github.com/morris-zheng/go-slim-micro-usersvc/internal/domain"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		server: grpc.NewServer(),
	}
}

func (s *GrpcServer) Run(ctx context.Context, svc *domain.ServiceContext) {
	name := "usersvc"
	if svc.Config.Name != "" {
		name = svc.Config.Name
	}

	var register *discovery.Register
	var err error
	etcdConfig := svc.Config.Etcd

	if len(etcdConfig.Endpoints) != 0 {
		svc.Logger.Info(ctx, "discovery register ...")
		register, err = discovery.NewRegister(discovery.Option{
			Endpoints: etcdConfig.Endpoints,
			Prefix:    etcdConfig.Prefix,
			TTL:       etcdConfig.KeepAliveInterval,
		}, &svc.Logger)

		if err != nil {
			svc.Logger.Fatal(ctx, "discovery register failed: %v", err)
		}

		err = register.Register(discovery.NewNode(name, svc.Config.Host, svc.Config.Port))
		if err != nil {
			svc.Logger.Fatal(ctx, "discovery register failed: %v", err)
		}
	}

	go func() {
		<-ctx.Done()
		if register != nil {
			svc.Logger.Info(context.Background(), "discovery deregister ...")
			register.Deregister()
		}

		s.server.GracefulStop()
		svc.Logger.Info(context.Background(), "shutdown server ...")
	}()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", svc.Config.Port))
	if err != nil {
		svc.Logger.Fatal(ctx, fmt.Sprintf("failed to listen: %v", err))
	}

	svc.Logger.Info(ctx, fmt.Sprintf("server listening at: %v", svc.Config.Port))
	if err := s.server.Serve(listen); err != nil {
		svc.Logger.Fatal(ctx, fmt.Sprintf(fmt.Sprintf("failed to serve: %v", err)))
	}
}

func (s *GrpcServer) Register(svc *domain.ServiceContext) {
	s.server.RegisterService(&user.Service_ServiceDesc, userDelivery.NewService(svc))
}
