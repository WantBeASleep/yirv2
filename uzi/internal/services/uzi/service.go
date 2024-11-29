package uzi

import (
	"context"
	"fmt"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateUzi(context.Context, domain.Uzi) (uuid.UUID, error)
	UpdateUzi(context.Context, uuid.UUID, OptionalUzi) (domain.Uzi, error)
	GetUzi(ctx context.Context, id uuid.UUID) (domain.Uzi, error)
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

func (s *service) CreateUzi(ctx context.Context, uzi domain.Uzi) (uuid.UUID, error) {
	uzi.Id = uuid.New()
	if err := s.dao.NewUziQuery(ctx).InsertUzi(uzi); err != nil {
		return uuid.Nil, fmt.Errorf("insert uzi: %w", err)
	}

	return uzi.Id, nil
}

func (s *service) UpdateUzi(ctx context.Context, id uuid.UUID, patch OptionalUzi) (domain.Uzi, error) {
	uzi, err := s.dao.NewUziQuery(ctx).UpdateUzi(id, patch.Map())
	if err != nil {
		return domain.Uzi{}, fmt.Errorf("update uzi: %w", err)
	}

	return uzi, nil
}

func (s *service) GetUzi(ctx context.Context, id uuid.UUID) (domain.Uzi, error) {
	uzi, err := s.dao.NewUziQuery(ctx).GetUziByPK(id)
	if err != nil {
		return domain.Uzi{}, fmt.Errorf("get uzi: %w", err)
	}

	return uzi, nil
}
