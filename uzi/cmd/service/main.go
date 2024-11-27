// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"log/slog"
	"net"
	"os"

	pkgconfig "yirv2/pkg/config"
	"yirv2/uzi/internal/config"

	"yirv2/uzi/internal/repository"

	uzisrv "yirv2/uzi/internal/services/uzi"

	pb "yirv2/uzi/internal/generated/grpc/service"
	grpchandler "yirv2/uzi/internal/grpc"
	devicehandler "yirv2/uzi/internal/grpc/device"
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
	slog.SetLogLoggerLevel(slog.LevelDebug)
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

	deviceHandler := devicehandler.New(dao)
	uziHandler := uzihandler.New(uziSrv)

	handler := grpchandler.New(
		deviceHandler,
		uziHandler,
	)

	server := grpc.NewServer()
	pb.RegisterUziSrvServer(server, handler)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	slog.Info("start serve")
	if err := server.Serve(lis); err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	return successExitCode
}
