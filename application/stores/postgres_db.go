package stores

import "database/sql"

type PostgresDb struct {
	Db *sql.DB
}
