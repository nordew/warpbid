package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	cockroach *sqlx.DB
	dragonfly *redis.Client
	sb        sq.StatementBuilderType
}

func New(primaryDB *sqlx.DB, cache *redis.Client) Storage {
	return Storage{
		cockroach: primaryDB,
		dragonfly: cache,
		sb:        sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
