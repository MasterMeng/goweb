package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// BlogDB is the interface that wrapper off SqlxDB
type BlogDB interface {
	IsInitialized() bool
	SetDBInitialized(bool)
	// BeginTx has same behavior as MustBegin except it returns FabricCATx
	// instead of *sqlx.Tx
	BeginTx() BlogTx
	DriverName() string

	Select(funcName string, dest interface{}, query string, args ...interface{}) error
	Exec(funcName, query string, args ...interface{}) (sql.Result, error)
	NamedExec(funcName, query string, arg interface{}) (sql.Result, error)
	Get(funcName string, dest interface{}, query string, args ...interface{}) error
	Queryx(funcName, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	MustBegin() *sqlx.Tx
	Close() error
	SetMaxOpenConns(n int)
	PingContext(ctx context.Context) error
}

// SqlxDB is the interface with functions implemented by sqlx.DB
// object that are used by Fabric CA server
type SqlxDB interface {
	DriverName() string
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	MustBegin() *sqlx.Tx
	Close() error
	SetMaxOpenConns(n int)
	PingContext(ctx context.Context) error
}

// DB is an adapter for sqlx.DB and implements BlogDB interface
type DB struct {
	DB SqlxDB
	// Indicates if database was successfully initialized
	IsDBInitialized bool
}

// New creates an instance of DB
func New(db SqlxDB) *DB {
	return &DB{
		DB: db,
	}
}

// IsInitialized returns true if db is initialized, else false
func (db *DB) IsInitialized() bool {
	return db.IsDBInitialized
}

// SetDBInitialized sets the value for Isdbinitialized
func (db *DB) SetDBInitialized(b bool) {
	db.IsDBInitialized = b
}

// BeginTx implements BeginTx method of BlogDB interface
func (db *DB) BeginTx() BlogTx {
	return &TX{
		TX: db.DB.MustBegin(),
	}
}

// Select performs select sql statement
func (db *DB) Select(funcName string, dest interface{}, query string, args ...interface{}) error {

	err := db.DB.Select(dest, query, args...)
	return err
}

// Exec executes query
func (db *DB) Exec(funcName, query string, args ...interface{}) (sql.Result, error) {

	res, err := db.DB.Exec(query, args...)

	return res, err
}

// NamedExec executes query
func (db *DB) NamedExec(funcName, query string, args interface{}) (sql.Result, error) {

	res, err := db.DB.NamedExec(query, args)
	return res, err
}

// Get executes query
func (db *DB) Get(funcName string, dest interface{}, query string, args ...interface{}) error {

	err := db.DB.Get(dest, query, args...)
	return err
}

// Queryx executes query
func (db *DB) Queryx(funcName, query string, args ...interface{}) (*sqlx.Rows, error) {

	rows, err := db.DB.Queryx(query, args...)
	return rows, err
}

// MustBegin starts a transaction
func (db *DB) MustBegin() *sqlx.Tx {
	return db.DB.MustBegin()
}

// DriverName returns database driver name
func (db *DB) DriverName() string {
	return db.DB.DriverName()
}

// Rebind parses query to properly format query
func (db *DB) Rebind(query string) string {
	return db.DB.Rebind(query)
}

// Close closes db
func (db *DB) Close() error {
	return db.DB.Close()
}

// SetMaxOpenConns sets number of max open connections
func (db *DB) SetMaxOpenConns(n int) {
	db.DB.SetMaxOpenConns(n)
}

// PingContext pings the database
func (db *DB) PingContext(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}
