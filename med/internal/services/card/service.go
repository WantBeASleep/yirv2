package card

import (
	"context"
	"fmt"

	"yirv2/med/internal/domain"
	"yirv2/med/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateCard(ctx context.Context, card domain.Card) error
	GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error)
	UpdateCard(ctx context.Context, card domain.Card) (domain.Card, error)
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

func (s *service) CreateCard(ctx context.Context, card domain.Card) error {
	if err := s.dao.NewCardQuery(ctx).InsertCard(card); err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}
// TODO: перенести все на entity
func (s *service) GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error) {
	card, err := s.dao.NewCardQuery(ctx).GetCardByPK(doctorID, patientID)
	if err != nil {
		return domain.Card{}, fmt.Errorf("insert card: %w", err)
	}

	return card, nil
}

// TODO: убрать это говно отсюда
func (s *service) UpdateCard(ctx context.Context, card domain.Card) (domain.Card, error) {
	card, err := s.dao.NewCardQuery(ctx).UpdateCard(card)
	if err != nil {
		return domain.Card{}, fmt.Errorf("insert card: %w", err)
	}

	return card, nil
}