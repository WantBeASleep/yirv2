package patient

import (
	"context"
	"fmt"

	"yirv2/med/internal/domain"
	"yirv2/med/internal/repository"
	"yirv2/med/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreatePatient(ctx context.Context, patient domain.Patient) (uuid.UUID, error)
	UpdatePatient(ctx context.Context, doctorID uuid.UUID, id uuid.UUID, path OptionalPatient) (domain.Patient, error)
	GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error)
	GetDoctorPatients(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) CreatePatient(ctx context.Context, patient domain.Patient) (uuid.UUID, error) {
	patient.Id = uuid.New()
	// TODO: нужно ввожить слой entity
	if err := s.dao.NewPatientQuery(ctx).InsertPatient(entity.Patient{}.FromDomain(patient)); err != nil {
		return uuid.Nil, fmt.Errorf("insert patient: %w", err)
	}

	return patient.Id, nil
}

func (s *service) UpdatePatient(
	ctx context.Context,
	doctorID uuid.UUID,
	id uuid.UUID,
	path OptionalPatient,
) (domain.Patient, error) {
	exists, err := s.dao.NewCardQuery(ctx).CheckCardExists(doctorID, id)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("check existing card")
	}
	if !exists {
		return domain.Patient{}, ErrNoPermission
	}

	patient, err := s.dao.NewPatientQuery(ctx).UpdatePatient(id, path.Map())
	if err != nil {
		return domain.Patient{}, fmt.Errorf("update patient: %w", err)
	}

	return patient, nil
}

func (s *service) GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error) {
	patient, err := s.dao.NewPatientQuery(ctx).GetPatientByPK(id)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("get patient: %w", err)
	}

	return patient, err
}

func (s *service) GetDoctorPatients(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error) {
	patients, err := s.dao.NewPatientQuery(ctx).GetPatientsByDoctorID(doctorID)
	if err != nil {
		return nil, fmt.Errorf("get doctor patients: %w", err)
	}

	return patients, err
}

