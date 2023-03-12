package ci

import (
	"log"
	"strings"
	"sync"
	"test/configs"
	"test/utils"

	"github.com/golang-migrate/migrate/v4"
)

var (
	once sync.Once
	cfg  = configs.Load()
)

func MigrationUp() {

	url, err := utils.ConnectionURLBuilder("migration")
	if err != nil {
		log.Println("Error generationg migration url", err.Error())
	}
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Fatal("error in creating migration: ", err.Error())
	}
	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Println("Erro in migrating ", err.Error())
		}
	}
}
