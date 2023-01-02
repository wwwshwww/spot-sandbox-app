package rdb

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/config"
	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	config.Configure()
}

func NewMySQLInstance(database string, tables ...any) (*gorm.DB, func() error, error) {
	cfg := mysql.Config{
		User:                 config.MySQL.User,
		Passwd:               config.MySQL.Passwd,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.MySQL.Host, config.MySQL.Port),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, nil, err
	}

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", database))
	if err != nil {
		return nil, nil, err
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database))
	if err != nil {
		return nil, nil, err
	}

	if err := db.Close(); err != nil {
		return nil, nil, err
	}

	cfg.DBName = database

	gormDB, err := gorm.Open(gormysql.Open(cfg.FormatDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, nil, err
	}
	if err := gormDB.AutoMigrate(tables...); err != nil {
		return nil, nil, err
	}

	db, err = gormDB.DB()
	if err != nil {
		return nil, nil, err
	}

	return gormDB, db.Close, nil
}
