package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var TestDB *sql.DB

func TestConnectDB() error {
	dsn := fmt.Sprintf(
		"%s@tcp(%s:%s)/%s?parseTime=true",
		"root",
		"phpmyadmin.test",
		"3306",
		"authify_db",
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	TestDB = db
	return nil
}
