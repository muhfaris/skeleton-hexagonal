package registry

import "database/sql"

type registry struct {
	db *sql.DB
}
