package mysql

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestConnect(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/test"
	db, err := sqlx.Open("mysql",dsn)
	if err != nil{
		t.Fatalf(err.Error())
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	res,err := db.Exec("CREATE TABLE test( id int);")
	if err != nil{
		t.Fatal(err.Error())
	}
	fmt.Println(res.RowsAffected())

	res,err = db.Exec("show tables;")
	if err != nil{
		t.Fatal(err.Error())
	}
	fmt.Println(res.RowsAffected())
}
