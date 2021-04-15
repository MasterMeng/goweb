package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"github.com/cloudflare/cfssl/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mastermeng/goweb/db"
	"github.com/mastermeng/goweb/db/util"
	"github.com/pkg/errors"
)

var (
	re = regexp.MustCompile(`\/([0-9,a-z,A-Z$_]+)`)
)

type Mysql struct {
	SqlxDB db.BlogDB

	datasource string
	dbName     string
}

func NewDB(datasource string) *Mysql {
	log.Debugf("Using MySQL database, connecting to database...")
	return &Mysql{
		datasource: datasource,
	}
}

func (m *Mysql) Connect() error {
	datasource := m.datasource

	m.dbName = util.GetDBName(datasource)
	log.Debugf("Database Name: %s", m.dbName)

	connStr := re.ReplaceAllString(datasource, "/")
	log.Debugf("Connecting to MySQL server, using connection string: %s", util.MaskDBCred(connStr))
	sqlxdb, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		return errors.Wrap(err, "Failed to connect to MySQL database")
	}

	m.SqlxDB = db.New(sqlxdb)
	return nil
}

// PingContext pings the database
func (m *Mysql) PingContext(ctx context.Context) error {
	err := m.SqlxDB.PingContext(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to ping to MySQL database")
	}
	return nil
}

// exists determines if the database has already been created
func exists(sqlxDB db.BlogDB, dbName string) (bool, error) {
	log.Debugf("Checking if MySQL CA database '%s' exists", dbName)
	var exists bool
	query := fmt.Sprintf("SELECT true as 'exists' FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s'", dbName)
	err := sqlxDB.Get("CheckIfDatabaseExists", &exists, query)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrapf(err, "Failed to check if MySQL Blog database '%s' exists", dbName)
	}
	return exists, nil
}

// Create creates database and tables
func (m *Mysql) Create() (*db.DB, error) {
	db, err := m.CreateDatabase()
	if err != nil {
		return nil, err
	}
	err = m.CreateTables()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateDatabase creates database
func (m *Mysql) CreateDatabase() (*db.DB, error) {
	err := m.createDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create MySQL database")
	}

	log.Debugf("Connecting to database '%s', using connection string: '%s'", m.dbName, util.MaskDBCred(m.datasource))
	sqlxdb, err := sqlx.Open("mysql", m.datasource)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to open database '%s' in MySQL server", m.dbName)
	}
	m.SqlxDB = db.New(sqlxdb)

	return m.SqlxDB.(*db.DB), nil
}

func (m *Mysql) createDatabase() error {
	exists, err := exists(m.SqlxDB, m.dbName)
	if err != nil {
		return err
	}
	if !exists {
		log.Debugf("Creating MySQL Database '%s'", m.dbName)
		_, err := m.SqlxDB.Exec("CreateDatabase", "CREATE DATABASE "+m.dbName)
		if err != nil {
			return errors.Wrap(err, "Failed to execute create database query")
		}
	}
	return nil
}

// CreateTables creates table
func (m *Mysql) CreateTables() error {
	err := m.createTables()
	if err != nil {
		return errors.Wrap(err, "Failed to create MySQL tables")
	}
	return nil
}

func (m *Mysql) createTables() error {
	db := m.SqlxDB
	log.Debug("Creating users table if it doesn`t exist")
	if _, err := db.Exec("CreateUsersTable", "CREATE TABLE IF NOT EXISTS users (id VARCHAR(255) NOT NULL, name VARCHAR(256), password VARCHAR(256), email VARCHAR(256), PRIMARY KEY (id)) DEFAULT CHARSET=utf8 COLLATE utf8_bin"); err != nil {
		return errors.Wrap(err, "Error creating users table")
	}

	log.Debug("Creating articles table if it doesn`t exist")
	if _, err := db.Exec("CreateArticlesTable", "CREATE TABLE IF NOT EXISTS articles (id VARCHAR(255) NOT NULL, owner VARCHAR(256) NOT NULL, title VARCHAR(256), type INTEGER, abstract TEXT, created INTEGER DEFAULT 0, modity INTEGER DEFAULT 0, content TEXT, PRIMARY KEY (id)) DEFAULT CHARSET=utf8 COLLATE utf8_bin"); err != nil {
		return errors.Wrap(err, "Error creating users table")
	}

	log.Debug("Creating messages table if it doesn`t exist")
	if _, err := db.Exec("CreateMessagesTable", "CREATE TABLE IF NOT EXISTS messages (id VARCHAR(255) NOT NULL, atricle VARCHAR(256) NOT NULL, user VARCHAR(256), created INTEGER DEFAULT 0, content TEXT,PRIMARY KEY (id)) DEFAULT CHARSET=utf8 COLLATE utf8_bin"); err != nil {
		return errors.Wrap(err, "Error creating users table")
	}

	return nil
}
