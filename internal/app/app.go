package app

import (
	"context"
	"log"
	"log/slog"
	"net"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/config"
	"github.com/pawpawchat/profile/internal/domain/service"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, config *config.Config) error {
	// листенер для сервера сервиса profile
	l, err := net.Listen("tcp", config.Env().GRPC_SERVER_ADDR)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	grpcServer := newGRPCServer(config)

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

func newGRPCServer(config *config.Config) *grpc.Server {
	db, err := sqlx.Connect("pgx", config.Env().DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	ps := service.NewProfile(repository.NewProfile(db))
	profileServer := newProfileGRPCServer(ps)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterProfileServiceServer(grpcServer, profileServer)
	return grpcServer
}
