package config

import (
	"fmt"
	"os"
)

type db struct {
	Driver       string
	rootPassword string
	user         string
	database     string
	hostWriter   string
	hostReader   string
	port         string
	charset      string
	parseTime    bool
}

func newDB() db {
	return db{
		Driver:       "mysql",
		rootPassword: os.Getenv("MYSQL_ROOT_PASSWORD"),
		user:         "root",
		database:     os.Getenv("MYSQL_DATABASE"),
		hostWriter:   os.Getenv("MYSQL_HOST_WRITER"),
		hostReader:   os.Getenv("MYSQL_HOST_READER"),
		port:         os.Getenv("MYSQL_PORT"),
		charset:      "utf8mb4",
		parseTime:    true,
	}
}

func (db db) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t",
		db.user, db.rootPassword, db.hostWriter, db.port, db.database, db.charset, db.parseTime,
	)
}
