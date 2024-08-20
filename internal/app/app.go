package app

import (
	"context"
	"log/slog"
	"net"
	"sync"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/config"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, config *config.Config) error {
	// листенер для сервера сервиса profile
	l, err := net.Listen("tcp", config.Env().GRPC_SERVER_ADDR)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	grpcServer := newGRPCServer()

	// запускаем сервер
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("grpc server is up and running", "addr", config.Env().GRPC_SERVER_ADDR)
		err = grpcServer.Serve(l)
	}()

	// запускаем горутину, которая ждет завершения контекста
	// и изящно останавливает запущенный сервер
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		grpcServer.GracefulStop()
		slog.Debug("grpc server was gracefuly stopped")
	}()

	wg.Wait()
	return err
}

func newGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	profileServer := NewServer()
	pb.RegisterProfileServiceServer(grpcServer, profileServer)
	return grpcServer
}
