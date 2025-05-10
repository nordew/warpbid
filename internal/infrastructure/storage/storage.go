package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	cockroach *sqlx.DB
	sb        sq.StatementBuilderType
}

func New(primaryDB *sqlx.DB) Storage {
	return Storage{
		cockroach: primaryDB,
		sb:        sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
