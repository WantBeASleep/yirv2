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
	runner Runner
}

func (q *BaseQuery) Context() context.Context {
	return q.ctx
}

func (q *BaseQuery) Runner() Runner {
	return q.runner
}

func (*BaseQuery) QueryBuilder() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}
