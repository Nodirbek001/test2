package postgres

import (
	"fmt"
	"sync"
	"test/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instanceG *gorm.DB
	once      sync.Once
)

func DB() *gorm.DB {
	once.Do(func() {
		dsn, err := utils.ConnectionURLBuilder("postgres")
		if err != nil {
			panic(err)
		}
		instanceG, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Gorm connections: %v", err.Error())
			panic(err)
		}
	})
	return instanceG
}


//go get -u
//go mod tidy
// go get -u ./...