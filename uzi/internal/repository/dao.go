package repository

import (
	"context"

	"yirv2/pkg/daolib"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewDeviceQuery(context.Context) DeviceQuery
	NewUziQuery(context.Context) UziQuery
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

func (d *dao) NewUziQuery(ctx context.Context) UziQuery {
	uziQuery := &uziQuery{}
	d.NewRepo(ctx, uziQuery)

	return uziQuery
}
