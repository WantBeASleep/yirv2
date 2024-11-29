package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	"yirv2/uzi/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const segmentTable = "segment"

type SegmentQuery interface {
	InsertSegment(segment domain.Segment) error
	GetSegmentByPK(id uuid.UUID) (domain.Segment, error)
	GetSegmentsByNodeID(id uuid.UUID) ([]domain.Segment, error)
	GetSegmentsByImageID(id uuid.UUID) ([]domain.Segment, error)
	UpdateSegment(id uuid.UUID, fields map[string]any) (domain.Segment, error)
	DeleteSegment(id uuid.UUID) error
	DeleteSegmentByUziID(id uuid.UUID) error
}

type segmentQuery struct {
	*daolib.BaseQuery
}

func (q *segmentQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *segmentQuery) InsertSegment(segment domain.Segment) error {
	query := q.QueryBuilder().
		Insert(segmentTable).
		Columns("id", "node_id", "image_id", "contor", "tirads_23", "tirads_4", "tirads_5").
		Values(segment.Id, segment.NodeID, segment.ImageID, segment.Contor,
			segment.Tirads23, segment.Tirads4, segment.Tirads5)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert segment: %w", err)
	}

	return nil
}

func (q *segmentQuery) GetSegmentByPK(id uuid.UUID) (domain.Segment, error) {
	query := q.QueryBuilder().
		Select("id", "node_id", "image_id", "contor", "tirads_23", "tirads_4", "tirads_5").
		From(segmentTable).
		Where(sq.Eq{
			"id": id,
		})

	var segments domain.Segment
	if err := q.Runner().Getx(q.Context(), &segments, query); err != nil {
		return domain.Segment{}, fmt.Errorf("get segments by pk: %w", err)
	}

	return segments, nil
}

func (q *segmentQuery) GetSegmentsByNodeID(id uuid.UUID) ([]domain.Segment, error) {
	query := q.QueryBuilder().
		Select("id", "node_id", "image_id", "contor", "tirads_23", "tirads_4", "tirads_5").
		From(segmentTable).
		Where(sq.Eq{
			"node_id": id,
		})

	var segments []domain.Segment
	if err := q.Runner().Selectx(q.Context(), &segments, query); err != nil {
		return nil, fmt.Errorf("get segments by uzi_id: %w", err)
	}

	return segments, nil
}

func (q *segmentQuery) GetSegmentsByImageID(id uuid.UUID) ([]domain.Segment, error) {
	query := q.QueryBuilder().
		Select("id", "node_id", "image_id", "contor", "tirads_23", "tirads_4", "tirads_5").
		From(segmentTable).
		Where(sq.Eq{
			"image_id": id,
		})

	var segments []domain.Segment
	if err := q.Runner().Selectx(q.Context(), &segments, query); err != nil {
		return nil, fmt.Errorf("get segments by image_id: %w", err)
	}

	return segments, nil
}

func (q *segmentQuery) UpdateSegment(id uuid.UUID, fields map[string]any) (domain.Segment, error) {
	query := q.QueryBuilder().
		Update(segmentTable).
		SetMap(fields).
		Where(sq.Eq{
			"id": id,
		}).
		Suffix("RETURNING *")

	var segment domain.Segment
	if err := q.Runner().Getx(q.Context(), &segment, query); err != nil {
		return domain.Segment{}, fmt.Errorf("update segment: %w", err)
	}

	return segment, nil
}

func (q *segmentQuery) DeleteSegment(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(segmentTable).
		Where(sq.Eq{
			"id": id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("delete segment: %w", err)
	}

	return nil
}

func (q *segmentQuery) DeleteSegmentByUziID(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(segmentTable).
		Where(sq.Eq{
			"node_id": id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("delete segment by uzi_id: %w", err)
	}

	return nil
}
