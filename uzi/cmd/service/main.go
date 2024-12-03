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

	devicesrv "yirv2/uzi/internal/services/device"
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
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", "err", err)
		return failExitCode
	}

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(db, client, "uzi")

	deviceSrv := devicesrv.New(dao)
	uziSrv := uzisrv.New(dao)
	imageSrv := imagesrv.New(dao)
	nodeSrv := nodesrv.New(dao)
	serviceSrv := segmentsrv.New(dao)

	deviceHandler := devicehandler.New(deviceSrv)
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
