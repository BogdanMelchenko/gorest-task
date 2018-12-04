package stores

import "database/sql"

type PostgresDbStore struct {
	Db *sql.DB
}
