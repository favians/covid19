package tests

import (
	"fmt"
	"os"

	"github.com/favians/golang_starter/bootstrap"
	"github.com/favians/golang_starter/db/migrations"
)

func init() {
	if bootstrap.App.ENV != "dev" && bootstrap.App.ENV != "test" {
		fmt.Println("Test only allowed on \"dev\" or \"test\" environment.")
		os.Exit(1)
	}
}

func RebuildData() {
	migrations.Truncate()
	migrations.Seed()
}
