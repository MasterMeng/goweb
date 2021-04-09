package factory

import (
	"context"

	"github.com/mastermeng/goweb/db"
	"github.com/mastermeng/goweb/db/mysql"
	"github.com/pkg/errors"
)

type DB interface {
	Connect() error
	PingContext(ctx context.Context) error
	Create() (*db.DB, error)
}

func New(dbType, datasource string) (DB, error) {
	switch dbType {
	case "mysql":
		return mysql.NewDB(datasource), nil
	default:
		return nil, errors.Errorf("Invalid db.type in config file: '%s'; must be 'sqlite3', 'postgres', or 'mysql'", dbType)
	}

}
