// Package database
// @author tabuyos
// @since 2023/6/30
// @description database
package database

import (
	"database/sql"
)

var db *sql.DB

func FetchDB() *sql.DB {
	return db
}
