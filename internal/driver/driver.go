package driver

import (
	"database/sql"
	"time"
)

// DB hold the connection
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenConns = 10
const maxIdleConns = 5
const maxLifeTime = 5 * time.Minute
