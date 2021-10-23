package database

import (
	"fmt"
	"sync"

	"github.com/jphacks/A_2108/src/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var (
	connectionPool map[string]*DatabaseHandler
)

type DatabaseHandler struct {
	DB   *gorm.DB
	lock *sync.Mutex
}

func getDSN(dbName string) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb",
		config.DBUser(),
		config.DBPass(),
		config.DBMethod(),
		dbName,
	)
}

func NewDatabaseHandlerWithDBName(dbName string) (*DatabaseHandler, error) {
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
	mut := &sync.Mutex{}
	res := &DatabaseHandler{
		DB:   db,
		lock: mut,
	}
	connectionPool[dbName] = res
	return res, nil
}

func NewDatabaseHandler() (*DatabaseHandler, error) {
	return NewDatabaseHandlerWithDBName("dbmaster")
}

func (h *DatabaseHandler) Lock() {
	h.lock.Lock()
}

func (h *DatabaseHandler) UnLock() {
	h.lock.Unlock()
}
