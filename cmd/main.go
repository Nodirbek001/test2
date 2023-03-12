package main

import (
	"test/ci"
	"test/configs"

	"github.com/blendle/zapdriver"

)

func main(){
	conf:=configs.Load()
	if err:=conf.Validate(); err!=nil {
		panic(err)
	}
	logger, err:=zapdriver.NewDevelopment()
	if err!=nil {
		panic(err)
	}
	ci.MigrationUp()


}