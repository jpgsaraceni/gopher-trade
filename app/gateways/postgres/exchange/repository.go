package exchange

import "github.com/jackc/pgx/v4/pgxpool"

const (
	uniqueKeyViolationCode = "23505"
	fromToDBConstraint     = "exchanges_from_to_key"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}
