package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(dialector gorm.Dialector) *gorm.DB {
	db := make(chan *gorm.DB)

	go func() {
		defer close(db)
		for {
			conn, err := gorm.Open(dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
			if err == nil {
				db <- conn
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()

	return <-db
}
