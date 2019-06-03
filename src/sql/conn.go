package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SQLdb struct {
	server   string
	user     string
	password string
	port     int
	dbName   string
	db       *sql.DB
}

func NewSQLdb(s string, u string, pw string, dbn string, p int) *SQLdb {
	return &SQLdb{
		server:   s,
		user:     u,
		password: pw,
		port:     p,
		dbName:   dbn,
	}
}

/*
type error interface {
    ToString() string
}
*/

func (s *SQLdb) Connect() error {
	dsn := fmt.Sprintf("%s:%s@/%s", s.user, s.password, s.dbName)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	s.db = conn
	return nil
}
