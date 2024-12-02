package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	"yirv2/uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const uziTable = "uzi"

type UziQuery interface {
	InsertUzi(uzi entity.Uzi) error
	GetUziByPK(id uuid.UUID) (entity.Uzi, error)
	UpdateUzi(uzi entity.Uzi) (int64, error)
}

type uziQuery struct {
	*daolib.BaseQuery
}

func (q *uziQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *uziQuery) InsertUzi(uzi entity.Uzi) error {
	query := q.QueryBuilder().
		Insert(uziTable).
		Columns(
			"id",
			"projection",
			"checked",
			"patient_id",
			"device_id",
			"create_at",
		).
		Values(
			uzi.Id,
			uzi.Projection,
			uzi.Checked,
			uzi.PatientID,
			uzi.DeviceID,
			uzi.CreateAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert uzi: %w", err)
	}

	return nil
}

func (q *uziQuery) GetUziByPK(id uuid.UUID) (entity.Uzi, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"projection",
			"checked",
			"patient_id",
			"device_id",
			"create_at",
		).
		From(uziTable).
		Where(sq.Eq{
			"id": id,
		})

	var uzi entity.Uzi
	if err := q.Runner().Getx(q.Context(), &uzi, query); err != nil {
		return entity.Uzi{}, fmt.Errorf("get uzi: %w", err)
	}

	return uzi, nil
}

func (q *uziQuery) UpdateUzi(uzi entity.Uzi) (int64, error) {
	query := q.QueryBuilder().
		Update(uziTable).
		SetMap(sq.Eq{
			"projection": uzi.Projection,
			"checked":    uzi.Checked,
		}).
		Where(sq.Eq{
			"id": uzi.Id,
		})

	rows, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, fmt.Errorf("update uzi: %w", err)
	}

	return rows.RowsAffected()
}
