package seeders

import (
	"github.com/Owoade/go-bank/sql"
)

type SeederStore struct {
	queries *sql.Queries
}

func NewSeedStore(q *sql.Queries) *SeederStore {
	return &SeederStore{
		queries: q,
	}
}
