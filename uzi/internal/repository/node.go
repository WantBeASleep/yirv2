package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	"yirv2/uzi/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const nodeTable = "node"

type NodeQuery interface {
	InsertNode(node domain.Node) error
	UpdateNode(id uuid.UUID, fields map[string]any) (domain.Node, error)
	GetNodesByImageID(id uuid.UUID) ([]domain.Node, error)
	DeleteNode(id uuid.UUID) error
}

type nodeQuery struct {
	*daolib.BaseQuery
}

func (q *nodeQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *nodeQuery) InsertNode(node domain.Node) error {
	query := q.QueryBuilder().
		Insert(nodeTable).
		Columns("id", "ai", "tirads_23", "tirads_4", "tirads_5").
		Values(node.Id, node.Ai, node.Tirads23, node.Tirads4, node.Tirads5)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert node: %w", err)
	}

	return nil
}

// TODO: упорядочнить Insert/Delete/Select/Update
func (q *nodeQuery) UpdateNode(id uuid.UUID, fields map[string]any) (domain.Node, error) {
	query := q.QueryBuilder().
		Update(nodeTable).
		SetMap(fields).
		Where(sq.Eq{
			"id": id,
		}).
		Suffix("RETURNING *")

	var uzi domain.Node
	if err := q.Runner().Getx(q.Context(), &uzi, query); err != nil {
		return domain.Node{}, fmt.Errorf("update node: %w", err)
	}

	return uzi, nil
}

func (q *nodeQuery) GetNodesByImageID(id uuid.UUID) ([]domain.Node, error) {
	query := q.QueryBuilder().
		Select("node.id", "node.ai", "node.tirads_23", "node.tirads_4", "node.tirads_5").
		From(nodeTable).
		InnerJoin("segment ON segment.node_id = node.id").
		InnerJoin("image ON image.id = segment.image_id").
		Where(sq.Eq{
			"image.id": id,
		})

	var uzi []domain.Node
	if err := q.Runner().Selectx(q.Context(), &uzi, query); err != nil {
		return nil, fmt.Errorf("get node by image_id: %w", err)
	}

	return uzi, nil
}

func (q *nodeQuery) DeleteNode(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(nodeTable).
		Where(sq.Eq{
			"id": id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("delete node: %w", err)
	}

	return nil
}
