package repository

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"yirv2/med/internal/domain"
	"yirv2/med/internal/repository/entity"
	"yirv2/pkg/daolib"
)

const patientTable = "patient"

type PatientQuery interface {
	InsertPatient(patient entity.Patient) error
	UpdatePatient(id uuid.UUID, fields map[string]any) (domain.Patient, error)
	GetPatientByPK(id uuid.UUID) (domain.Patient, error)
	GetPatientsByDoctorID(id uuid.UUID) ([]domain.Patient, error)
}

type patientQuery struct {
	*daolib.BaseQuery
}

func (q *patientQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *patientQuery) InsertPatient(patient entity.Patient) error {
	query := q.QueryBuilder().
		Insert(patientTable).
		Columns(
			"id",
			"fullname",
			"email",
			"policy",
			"active",
			"malignancy",
			"lastUziDate",
		).
		Values(
			patient.Id,
			patient.FullName,
			patient.Email,
			patient.Policy,
			patient.Active,
			patient.Malignancy,
			patient.LastUziDate,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert patient: %w", err)
	}

	return nil
}

func (q *patientQuery) UpdatePatient(id uuid.UUID, fields map[string]any) (domain.Patient, error) {
	query := q.QueryBuilder().
		Update(patientTable).
		SetMap(fields).
		Where(sq.Eq{
			"id": id,
		}).
		Suffix("RETURNING *")

	var patient domain.Patient
	if err := q.Runner().Getx(q.Context(), &patient, query); err != nil {
		return domain.Patient{}, fmt.Errorf("update patient: %w", err)
	}

	return patient, nil
}

func (q *patientQuery) GetPatientByPK(id uuid.UUID) (domain.Patient, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"fullname",
			"email",
			"policy",
			"active",
			"malignancy",
			"lastUziDate",
		).
		From(patientTable).
		Where(sq.Eq{
			"id": id,
		})

	var patient domain.Patient
	if err := q.Runner().Getx(q.Context(), &patient, query); err != nil {
		return domain.Patient{}, fmt.Errorf("get patient: %w", err)
	}

	return patient, nil
}

func (q *patientQuery) GetPatientsByDoctorID(id uuid.UUID) ([]domain.Patient, error) {
	query := q.QueryBuilder().
		Select(
			"patient.id",
			"patient.fullname",
			"patient.email",
			"patient.policy",
			"patient.active",
			"patient.malignancy",
			"patient.lastUziDate",
		).
		From(patientTable).
		InnerJoin("card ON card.patient_id = patient.id").
		Where(sq.Eq{
			"card.doctor_id": id,
		})

	var patient []domain.Patient
	if err := q.Runner().Selectx(q.Context(), &patient, query); err != nil {
		return nil, fmt.Errorf("get patient: %w", err)
	}

	return patient, nil
}