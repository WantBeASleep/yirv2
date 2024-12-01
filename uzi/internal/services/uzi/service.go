package uzi

import (
	"context"
	"errors"
	"fmt"
	"time"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateUzi(context.Context, domain.Uzi) (uuid.UUID, error)
	UpdateUzi(context.Context, uuid.UUID, OptionalUzi) (domain.Uzi, error)
	GetUzi(ctx context.Context, id uuid.UUID) (domain.Uzi, domain.Echographic, error)
	UpdateEchographic(ctx context.Context, id uuid.UUID, patch OptionalEchographic) (domain.Echographic, error)
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
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("start transaction: %w", err)
	}

	uzi.Id = uuid.New()
	uzi.Checked = false
	uzi.CreateAt = time.Now()

	if err := s.dao.NewUziQuery(ctx).InsertUzi(uzi); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert uzi: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.NewEchographicQuery(ctx).InsertEchographic(domain.Echographic{Id: uzi.Id}); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert uzi: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
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

func (s *service) GetUzi(ctx context.Context, id uuid.UUID) (domain.Uzi, domain.Echographic, error) {
	uzi, err := s.dao.NewUziQuery(ctx).GetUziByPK(id)
	if err != nil {
		return domain.Uzi{}, domain.Echographic{}, fmt.Errorf("get uzi: %w", err)
	}

	echographic, err := s.dao.NewEchographicQuery(ctx).GetEchographicByPK(id)
	if err != nil {
		return domain.Uzi{}, domain.Echographic{}, fmt.Errorf("get echographic: %w", err)
	}

	return uzi, echographic, nil
}

func (s *service) UpdateEchographic(ctx context.Context, id uuid.UUID, patch OptionalEchographic) (domain.Echographic, error) {
	echographic, err := s.dao.NewEchographicQuery(ctx).UpdateEchographic(id, patch.Map())
	if err != nil {
		return domain.Echographic{}, fmt.Errorf("update echographic: %w", err)
	}

	return echographic, nil
}
