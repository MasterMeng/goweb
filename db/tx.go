package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// BlogTx is the interface with functions implemented by sqlx.Tx
// object that are used by Fabric CA server
type BlogTx interface {
	Select(funcName string, dest interface{}, query string, args ...interface{}) error
	Exec(funcName, query string, args ...interface{}) (sql.Result, error)
	Queryx(funcName, query string, args ...interface{}) (*sqlx.Rows, error)
	Get(funcName string, dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
	Commit(funcName string) error
	Rollback(funcName string) error
}

// SqlxTx is the contract with sqlx
type SqlxTx interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
	Exec(query string, args ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}

// TX is the database transaction
type TX struct {
	TX SqlxTx
}

// Select performs select sql statement
func (tx *TX) Select(funcName string, dest interface{}, query string, args ...interface{}) error {
	err := tx.TX.Select(dest, query, args...)
	return err
}

// Exec executes query
func (tx *TX) Exec(funcName, query string, args ...interface{}) (sql.Result, error) {
	res, err := tx.TX.Exec(query, args...)
	return res, err
}

// Get executes query
func (tx *TX) Get(funcName string, dest interface{}, query string, args ...interface{}) error {

	err := tx.TX.Get(dest, query, args...)
	return err
}

// Queryx executes query
func (tx *TX) Queryx(funcName, query string, args ...interface{}) (*sqlx.Rows, error) {

	rows, err := tx.TX.Queryx(query, args...)
	return rows, err
}

// Rebind rebinds the query
func (tx *TX) Rebind(query string) string {
	return tx.TX.Rebind(query)
}

// Commit commits the transaction
func (tx *TX) Commit(funcName string) error {
	err := tx.TX.Commit()
	return err
}

// Rollback roll backs the transaction
func (tx *TX) Rollback(funcName string) error {
	err := tx.TX.Rollback()
	return err
}
