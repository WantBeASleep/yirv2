package image

import (
	"context"
	"fmt"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	GetUziImages(ctx context.Context, uziID uuid.UUID) ([]domain.Image, error)
	GetImageSegmentsWithNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error)
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

func (s *service) GetUziImages(ctx context.Context, uziID uuid.UUID) ([]domain.Image, error) {
	images, err := s.dao.NewImageQuery(ctx).GetImagesByUziID(uziID)
	if err != nil {
		return nil, fmt.Errorf("get images by uzi_id: %w", err)
	}

	return images, nil
}

func (s *service) GetImageSegmentsWithNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error) {
	segments, err := s.dao.NewSegmentQuery(ctx).GetSegmentsByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get segments by image_id: %w", err)
	}

	// TODO: переделать на запросе без JOIN
	nodes, err := s.dao.NewNodeQuery(ctx).GetNodesByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get nodes by image_id: %w", err)
	}

	return nodes, segments, nil
}
