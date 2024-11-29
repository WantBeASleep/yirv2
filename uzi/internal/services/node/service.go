package node

import (
	"context"
	"errors"
	"fmt"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateNode(ctx context.Context, node domain.Node, segments []domain.Segment) (uuid.UUID, error)
	UpdateNode(ctx context.Context, id uuid.UUID, patch OptionalNode) (domain.Node, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
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

// TODO: rollback через defer
func (s *service) CreateNode(ctx context.Context, node domain.Node, segments []domain.Segment) (uuid.UUID, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("start transaction: %w", err)
	}
	defer func() {}()

	node.Id = uuid.New()
	// ai create через брокер
	node.Ai = false
	if err := s.dao.NewNodeQuery(ctx).InsertNode(node); err != nil {
		// TODO: rollback нормально ошибку оформить
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert node: %w", errors.Join(err, rollbackErr))
	}

	segmentQuery := s.dao.NewSegmentQuery(ctx)
	for _, v := range segments {
		v.Id = uuid.New()
		v.NodeID = node.Id

		if err := segmentQuery.InsertSegment(v); err != nil {
			rollbackErr := s.dao.RollbackTx(ctx)
			return uuid.Nil, fmt.Errorf("insert segment: %w", errors.Join(err, rollbackErr))
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
	}

	return node.Id, nil
}

func (s *service) DeleteNode(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}

	if err := s.dao.NewSegmentQuery(ctx).DeleteSegmentByUziID(id); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("delete node segments: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.NewNodeQuery(ctx).DeleteNode(id); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("delete node: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (s *service) UpdateNode(ctx context.Context, id uuid.UUID, patch OptionalNode) (domain.Node, error) {
	node, err := s.dao.NewNodeQuery(ctx).UpdateNode(id, patch.Map())
	if err != nil {
		return domain.Node{}, fmt.Errorf("update node: %w", err)
	}

	return node, nil
}
