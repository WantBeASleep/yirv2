package grpc

import (
	"yirv2/uzi/internal/generated/grpc/service"
	"yirv2/uzi/internal/grpc/device"
	"yirv2/uzi/internal/grpc/uzi"
)

type Handler struct {
	device.DeviceHandler
	uzi.UziHandler

	service.UnsafeUziSrvServer
}

func New(
	deviceHandler device.DeviceHandler,
	uziHandler uzi.UziHandler,
) *Handler {
	return &Handler{
		DeviceHandler: deviceHandler,
		UziHandler:    uziHandler,
	}
}
