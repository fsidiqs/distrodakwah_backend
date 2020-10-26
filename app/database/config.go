package database

import (
	"fmt"
	"os"
)

func BuildDSN() string {
	var dsn string
	env := os.Getenv("ENV")
	_ = env
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	user := os.Getenv("DB_USERNAME")
	if user == "" {
		user = "root"
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = "root"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "dd"
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return dsn
}
