package main

import (
	"fmt"
	"github.com/anujk14/go-playground/internal"
	"os"

	"github.com/anujk14/go-playground/internal/server"
	"github.com/anujk14/go-playground/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger := logger.GetLogger()

	config := internal.GetConfig()
	logger.Info(fmt.Sprintf("Database URL: %s", config.DatabaseURL))

	args := os.Args
	if len(args) > 1 {
		if args[1] == "migrate" {
			m, err := migrate.New("file://db/migrations", config.DatabaseURL)
			if err != nil {
				logger.Error("Error in initializing migrations: ", err.Error())
				os.Exit(1)
			}

			if err := m.Up(); err != nil {
				logger.Error("Error in running migrations: ", err.Error())
			}
		} else {
			fmt.Println("Not a valid command")
		}
	} else {
		server.Init()
	}
}
