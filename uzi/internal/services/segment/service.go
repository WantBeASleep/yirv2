package segment

import (
	"context"
	"errors"
	"fmt"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateSegment(ctx context.Context, segment domain.Segment) (uuid.UUID, error)
	UpdateSegment(ctx context.Context, id uuid.UUID, patch OptionalSegment) (domain.Segment, error)
	DeleteSegment(ctx context.Context, id uuid.UUID) error
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

func (s *service) CreateSegment(ctx context.Context, segment domain.Segment) (uuid.UUID, error) {
	segment.Id = uuid.New()
	if err := s.dao.NewSegmentQuery(ctx).InsertSegment(segment); err != nil {
		return uuid.Nil, fmt.Errorf("insert segment: %w", err)
	}

	return segment.Id, nil
}

func (s *service) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}

	segmentQuery := s.dao.NewSegmentQuery(ctx)

	segment, err := segmentQuery.GetSegmentByPK(id)
	if err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("get segment: %w", errors.Join(err, rollbackErr))
	}

	if err := segmentQuery.DeleteSegment(id); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("delete segment: %w", errors.Join(err, rollbackErr))
	}

	remainingSegments, err := segmentQuery.GetSegmentsByNodeID(segment.NodeID)
	if err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("get segment by uzi_id: %w", errors.Join(err, rollbackErr))
	}

	// у node не осталось сегментов, удаляем
	if len(remainingSegments) == 0 {
		if err := s.dao.NewNodeQuery(ctx).DeleteNode(segment.NodeID); err != nil {
			rollbackErr := s.dao.RollbackTx(ctx)
			return fmt.Errorf("delete node wo segments: %w", errors.Join(err, rollbackErr))
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (s *service) UpdateSegment(ctx context.Context, id uuid.UUID, patch OptionalSegment) (domain.Segment, error) {
	segment, err := s.dao.NewSegmentQuery(ctx).UpdateSegment(id, patch.Map())
	if err != nil {
		return domain.Segment{}, fmt.Errorf("update segment: %w", err)
	}

	return segment, nil
}
