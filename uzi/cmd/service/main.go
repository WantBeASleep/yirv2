// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"log/slog"
	"net"
	"os"

	pkgconfig "yirv2/pkg/config"
	"yirv2/pkg/grpclib"
	"yirv2/pkg/loglib"
	"yirv2/uzi/internal/config"

	"yirv2/uzi/internal/repository"

	imagesrv "yirv2/uzi/internal/services/image"
	nodesrv "yirv2/uzi/internal/services/node"
	segmentsrv "yirv2/uzi/internal/services/segment"
	uzisrv "yirv2/uzi/internal/services/uzi"

	pb "yirv2/uzi/internal/generated/grpc/service"
	grpchandler "yirv2/uzi/internal/grpc"

	devicehandler "yirv2/uzi/internal/grpc/device"
	imagehandler "yirv2/uzi/internal/grpc/image"
	nodehandler "yirv2/uzi/internal/grpc/node"
	segmenthandler "yirv2/uzi/internal/grpc/segment"
	uzihandler "yirv2/uzi/internal/grpc/uzi"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	defaultCfgPath = "service.yml"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	loglib.InitLogger(loglib.WithDevEnv())
	cfg, err := pkgconfig.Load[config.Config](defaultCfgPath)
	if err != nil {
		slog.Error("init config", "err", err)
		return failExitCode
	}

	db, err := sqlx.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		slog.Error("init db", "err", err)
		return failExitCode
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(db)

	uziSrv := uzisrv.New(dao)
	imageSrv := imagesrv.New(dao)
	nodeSrv := nodesrv.New(dao)
	serviceSrv := segmentsrv.New(dao)

	deviceHandler := devicehandler.New(dao)
	uziHandler := uzihandler.New(uziSrv)
	imageHandler := imagehandler.New(imageSrv)
	nodeHandler := nodehandler.New(nodeSrv)
	serviceHandler := segmenthandler.New(serviceSrv)

	handler := grpchandler.New(
		deviceHandler,
		uziHandler,
		imageHandler,
		nodeHandler,
		serviceHandler,
	)

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(grpclib.ServerCallLoggerInterceptor))
	pb.RegisterUziSrvServer(server, handler)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	if err := server.Serve(lis); err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	return successExitCode
}
