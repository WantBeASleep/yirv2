package repository

import (
	"context"

	"yirv2/pkg/daolib"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewDeviceQuery(ctx context.Context) DeviceQuery
	NewUziQuery(ctx context.Context) UziQuery
	NewImageQuery(ctx context.Context) ImageQuery
	NewSegmentQuery(ctx context.Context) SegmentQuery
	NewNodeQuery(ctx context.Context) NodeQuery
	NewEchographicQuery(ctx context.Context) EchographicQuery
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

func (d *dao) NewImageQuery(ctx context.Context) ImageQuery {
	imageQuery := &imageQuery{}
	d.NewRepo(ctx, imageQuery)

	return imageQuery
}

func (d *dao) NewSegmentQuery(ctx context.Context) SegmentQuery {
	segment := &segmentQuery{}
	d.NewRepo(ctx, segment)

	return segment
}

func (d *dao) NewNodeQuery(ctx context.Context) NodeQuery {
	node := &nodeQuery{}
	d.NewRepo(ctx, node)

	return node
}

func (d *dao) NewEchographicQuery(ctx context.Context) EchographicQuery {
	echographic := &echographicQuery{}
	d.NewRepo(ctx, echographic)

	return echographic
}
