package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	"yirv2/uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const imageTable = "image"

type ImageQuery interface {
	GetImagesByUziID(uziID uuid.UUID) ([]entity.Image, error)
}

type imageQuery struct {
	*daolib.BaseQuery
}

func (q *imageQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *imageQuery) GetImagesByUziID(uziID uuid.UUID) ([]entity.Image, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"page",
		).
		From(imageTable).
		Where(sq.Eq{
			"uzi_id": uziID,
		})

	var images []entity.Image
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, fmt.Errorf("get image by uzi_id: %w", err)
	}

	return images, nil
}
