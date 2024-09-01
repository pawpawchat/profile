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
	"github.com/pawpawchat/profile/internal/app/grpc/server"
	"github.com/pawpawchat/profile/internal/domain/service/avatar"
	"github.com/pawpawchat/profile/internal/domain/service/profile"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"github.com/pawpawchat/profile/pkg/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, env config.Environment) {
	l, err := net.Listen("tcp", env.GRPC_SERVER_ADDR)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	grpcServer := newGRPCServer(env)

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("grpc server is up and running", "addr", env.GRPC_SERVER_ADDR)
		if err := grpcServer.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		grpcServer.GracefulStop()
		slog.Debug("grpc server was gracefuly stopped")
	}()

	wg.Wait()
}

func newGRPCServer(env config.Environment) *grpc.Server {
	pr, br, ar := createRepositoriesWithOneConn(env.DB_URL)

	ps := profile.NewProfileService(pr, br)
	as := avatar.NewAvatarService(ar)

	profileServer := server.NewProfileGRPCServer(ps, as)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryInterceptor))

	pb.RegisterProfileServiceServer(grpcServer, profileServer)

	reflection.Register(grpcServer)
	return grpcServer
}

func createRepositoriesWithOneConn(url string) (*repository.ProfileRepository, *repository.BiographyRepository, *repository.AvatarRepository) {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatal(err)
	}

	pr := repository.NewProfileRepository(db)
	br := repository.NewBiographyRepository(db)
	ar := repository.NewAvatarsRepository(db)

	return pr, br, ar
}
