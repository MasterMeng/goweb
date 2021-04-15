package mysql

import (
	"context"
	"testing"
)

func TestConnect(t *testing.T) {
	dsn := "root:^n~GXbxv#gGm[&O`[o8Z@tcp(localhost:3306)/blog"

	mysql := NewDB(dsn)
	err := mysql.Connect()
	if err != nil {
		panic(err.Error())
	}

	err = mysql.PingContext(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	_, err = mysql.CreateDatabase()
	if err != nil {
		panic(err.Error())
	}
	if err := mysql.CreateTables(); err != nil {
		panic(err.Error())
	}
}
