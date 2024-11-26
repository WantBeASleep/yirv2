package repository

import (
	"context"

	"yirv2/pkg/daolib"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewDeviceQuery(ctx context.Context) DeviceQuery
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewDeviceQuery(ctx context.Context) DeviceQuery {
	deviceQuery := &deviceQuery{}
	d.NewRepo(ctx, deviceQuery)

	return deviceQuery
}
