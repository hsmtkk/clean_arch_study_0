package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hsmtkk/clean_arch_study_0/pkg/interface/database"
)

type sqlHandler struct {
	conn *sql.DB
}

func New() (database.SQLHandler, error) {
	conn, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return nil, err
	}
	return &sqlHandler{conn: conn}, nil
}

func (handler *sqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := handler.conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.result = result
	return res, nil
}

func (handler *sqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	row := SQLRow{}
	rows, err := handler.conn.Query(statement, args...)
	if err != nil {
		return row, err
	}
	row.rows = rows
	return row, nil
}

type SQLResult struct {
	result sql.Result
}

func (r SQLResult) LastInsertID() (int64, error) {
	return r.result.LastInsertId()
}

func (r SQLResult) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}

type SQLRow struct {
	rows *sql.Rows
}

func (r SQLRow) Scan(dest ...interface{}) error {
	return r.rows.Scan()
}

func (r SQLRow) Next() bool {
	return r.rows.Next()
}

func (r SQLRow) Close() error {
	return r.rows.Close()
}
