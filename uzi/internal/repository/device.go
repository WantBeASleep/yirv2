package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	"yirv2/uzi/internal/domain"
)

const deviceTable = "device"

type DeviceQuery interface {
	GetDeviceList() ([]domain.Device, error)
}

type deviceQuery struct {
	*daolib.BaseQuery
}

func (q *deviceQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *deviceQuery) GetDeviceList() ([]domain.Device, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
		).
		From(deviceTable)

	var devices []domain.Device
	if err := q.Runner().Selectx(q.Context(), &devices, query); err != nil {
		return nil, fmt.Errorf("get device list: %w", err)
	}

	return devices, nil
}
