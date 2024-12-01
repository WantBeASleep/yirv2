package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"yirv2/med/internal/domain"
)

const doctorTable = "doctor"

type DoctorQuery interface {
	InsertDoctor(doctor domain.Doctor) error
	UpdateUzi(id uuid.UUID, fields map[string]any) (domain.Doctor, error)
	GetDoctorByPK(id uuid.UUID) (domain.Doctor, error)
}

type doctorQuery struct {
	*daolib.BaseQuery
}

func (q *doctorQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *doctorQuery) InsertDoctor(doctor domain.Doctor) error {
	query := q.QueryBuilder().
		Insert(doctorTable).
		Columns("id", "fullname", "org", "job", "desc").
		Values(doctor.Id, doctor.FullName, doctor.Org, doctor.Job, doctor.Desc)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert doctor: %w", err)
	}

	return nil
}

func (q *doctorQuery) UpdateUzi(id uuid.UUID, fields map[string]any) (domain.Doctor, error) {
	query := q.QueryBuilder().
		Update(doctorTable).
		SetMap(fields).
		Where(sq.Eq{
			"id": id,
		}).
		Suffix("RETURNING *")

	var doctor domain.Doctor
	if err := q.Runner().Getx(q.Context(), &doctor, query); err != nil {
		return domain.Doctor{}, fmt.Errorf("update doctor: %w", err)
	}

	return doctor, nil
}

func (q *doctorQuery) GetDoctorByPK(id uuid.UUID) (domain.Doctor, error) {
	query := q.QueryBuilder().
		Select("id", "fullname", "org", "job", "desc").
		From(doctorTable).
		Where(sq.Eq{
			"id": id,
		})

	var doctor domain.Doctor
	if err := q.Runner().Getx(q.Context(), &doctor, query); err != nil {
		return domain.Doctor{}, fmt.Errorf("get doctor: %w", err)
	}

	return doctor, nil
}