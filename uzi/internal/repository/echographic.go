package repository

import (
	"fmt"

	"yirv2/pkg/daolib"
	"yirv2/uzi/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const echographicTable = "echographic"

type EchographicQuery interface {
	InsertEchographic(echographic domain.Echographic) error
	UpdateEchographic(id uuid.UUID, fields map[string]any) (domain.Echographic, error)
	GetEchographicByPK(id uuid.UUID) (domain.Echographic, error)
}

type echographicQuery struct {
	*daolib.BaseQuery
}

func (q *echographicQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *echographicQuery) InsertEchographic(echographic domain.Echographic) error {
	query := q.QueryBuilder().
		Insert(echographicTable).
		Columns(
			"id",
			"contors",
			"left_lobe_length",
			"left_lobe_width",
			"left_lobe_thick",
			"left_lobe_volum",
			"right_lobe_length",
			"right_lobe_width",
			"right_lobe_thick",
			"right_lobe_volum",
			"gland_volum",
			"isthmus",
			"struct",
			"echogenicity",
			"regional_lymph",
			"vascularization",
			"location",
			"additional",
			"conclusion",
		).
		Values(
			echographic.Id,
			echographic.Contors,
			echographic.LeftLobeLength,
			echographic.LeftLobeWidth,
			echographic.LeftLobeThick,
			echographic.LeftLobeVolum,
			echographic.RightLobeLength,
			echographic.RightLobeWidth,
			echographic.RightLobeThick,
			echographic.RightLobeVolum,
			echographic.GlandVolum,
			echographic.Isthmus,
			echographic.Struct,
			echographic.Echogenicity,
			echographic.RegionalLymph,
			echographic.Vascularization,
			echographic.Location,
			echographic.Additional,
			echographic.Conclusion,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert echographic: %w", err)
	}

	return nil
}

func (q *echographicQuery) UpdateEchographic(id uuid.UUID, fields map[string]any) (domain.Echographic, error) {
	query := q.QueryBuilder().
		Update(echographicTable).
		SetMap(fields).
		Where(sq.Eq{
			"id": id,
		}).
		Suffix("RETURNING *")

	var echographic domain.Echographic
	if err := q.Runner().Getx(q.Context(), &echographic, query); err != nil {
		return domain.Echographic{}, fmt.Errorf("update echographic: %w", err)
	}

	return echographic, nil
}

func (q *echographicQuery) GetEchographicByPK(id uuid.UUID) (domain.Echographic, error) {
	query := q.QueryBuilder().
		Select("*").
		From(echographicTable).
		Where(sq.Eq{
			"id": id,
		})

	var echographic domain.Echographic
	if err := q.Runner().Getx(q.Context(), &echographic, query); err != nil {
		return domain.Echographic{}, fmt.Errorf("get echographic: %w", err)
	}

	return echographic, nil
}
