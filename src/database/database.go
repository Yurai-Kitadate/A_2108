package database

import (
	"fmt"

	"github.com/jphacks/A_2108/src/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var (
	connectionPool map[string]*gorm.DB
)

func getDSN(dbName string) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=true",
		config.DBUser(),
		config.DBPass(),
		config.DBMethod(),
		dbName,
	)
}

func NewDatabaseHandlerWithDBName(dbName string) (*gorm.DB, error) {
	if connectionPool == nil {
		connectionPool = map[string]*gorm.DB{}
	}

	{
		has, ok := connectionPool[dbName]
		if ok {
			return has, nil
		}
	}
	dsn := getDSN(dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	connectionPool[dbName] = db
	return db, nil
}

func NewDatabaseHandler() (*gorm.DB, error) {
	return NewDatabaseHandlerWithDBName("dbmaster")
}
