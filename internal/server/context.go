package server

import "database/sql"

type Context struct {
	Db *sql.DB
}
