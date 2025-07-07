package dbconnection

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
