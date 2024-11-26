// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"log/slog"
	"net"
	"os"

	pkgconfig "yirv2/pkg/config"
	"yirv2/uzi/internal/config"
	pb "yirv2/uzi/internal/generated/grpc/service"
	devicehandler "yirv2/uzi/internal/grpc/device"
	"yirv2/uzi/internal/repository"
	devicesrv "yirv2/uzi/internal/services/device"

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
	deviceSrv := devicesrv.NewDeviceService(dao)
	deviceHandler := devicehandler.NewHandler(deviceSrv)

	server := grpc.NewServer()
	pb.RegisterDeviceSrvServer(server, deviceHandler)

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
