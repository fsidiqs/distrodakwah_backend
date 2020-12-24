package database

import (
	"fmt"
	"os"
)

func BuildDSN() string {
	var dsn string
	var host, port, user, pass, dbname string

	env := os.Getenv("ENV")
	if env == "STAGING" {
		host = os.Getenv("DB_HOST_STAGING")
		// port = os.Getenv("DB_PORT_STAGING")
		user = os.Getenv("DB_USERNAME_STAGING")
		pass = os.Getenv("DB_PASSWORD_STAGING")
		dbname = os.Getenv("DB_NAME_STAGING")
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, pass, host, dbname)

	}
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "3306"
	}

	if user == "" {
		user = "root"
	}

	if pass == "" {
		pass = os.Getenv("DB_PASS")
	}

	if dbname == "" {
		dbname = os.Getenv("DB_NAME")
	}

	if env != "STAGING" || env != "PRODUCTION" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, pass, host, port, dbname)
	}
	return dsn
}
