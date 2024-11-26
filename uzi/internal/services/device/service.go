package device

import (
	"context"
	"fmt"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"
)

type Device interface {
	GetDeviceList(ctx context.Context) ([]domain.Device, error)
}

type deviceService struct {
	dao repository.DAO
}

func NewDeviceService(
	dao repository.DAO,
) Device {
	return &deviceService{
		dao: dao,
	}
}

func (s *deviceService) GetDeviceList(ctx context.Context) ([]domain.Device, error) {
	devices, err := s.dao.NewDeviceQuery(ctx).GetDeviceList()
	if err != nil {
		return nil, fmt.Errorf("get device list from repo: %w", err)
	}
	return devices, nil
}
