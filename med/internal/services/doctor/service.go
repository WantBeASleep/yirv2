package doctor

import (
	"context"
	"fmt"

	"yirv2/med/internal/domain"
	"yirv2/med/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	RegisterDoctor(ctx context.Context, doctor domain.Doctor) error
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

func (s *service) RegisterDoctor(ctx context.Context, doctor domain.Doctor) error {
	if err := s.dao.NewDoctorQuery(ctx).InsertDoctor(doctor); err != nil {
		return fmt.Errorf("insert doctor: %w", err)
	}

	return nil
}

func (s *service) UpdateDoctor(ctx context.Context, id uuid.UUID, patch OptionalDoctor) (domain.Doctor, error) {
	doctor, err := s.dao.NewDoctorQuery(ctx).UpdateUzi(id, patch.Map())
	if err != nil {
		return domain.Doctor{}, fmt.Errorf("update doctor: %w", err)
	}

	return doctor, nil
}

func (s *service) GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error) {
	doctor, err := s.dao.NewDoctorQuery(ctx).GetDoctorByPK(id)
	if err != nil {
		return domain.Doctor{}, fmt.Errorf("update doctor: %w", err)
	}

	return doctor, nil
}