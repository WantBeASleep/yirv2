package gtclib

import (
	"database/sql"
	"time"
)

type _time struct{}

var Time _time

func (_time) PointerToSql(p *time.Time) sql.NullTime {
	if p == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *p, Valid: true}
}