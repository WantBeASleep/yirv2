package device

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "yirv2/uzi/internal/generated/grpc/service"
	"yirv2/uzi/internal/services/device"
)

type Handler struct {
	deviceSrv device.Device

	pb.UnimplementedDeviceSrvServer
}

func NewHandler(
	deviceSrv device.Device,
) *Handler {
	return &Handler{
		deviceSrv: deviceSrv,
	}
}

func (h *Handler) GetDeviceList(ctx context.Context, _ *empty.Empty) (*pb.GetDeviceListOut, error) {
	devices, err := h.deviceSrv.GetDeviceList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetDeviceListOut)
	for _, d := range devices {
		pbDevice := domainDeviceToPbDevice(&d)
		out.Devices = append(out.Devices, &pbDevice)
	}

	return out, nil
}
