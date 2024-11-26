package daolib

import (
	"context"

	"github.com/Masterminds/squirrel"
)

type BaseQuerySetter interface {
	SetBaseQuery(*BaseQuery)
}

type BaseQuery struct {
	ctx context.Context
	Runner
}

func (BaseQuery) QueryBuilder() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}
