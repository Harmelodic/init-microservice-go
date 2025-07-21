package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
)

type appConfig struct {
	MigrationsDirectory string
	DbConnectionString  string
}

func loadAppConfigFromCommandFlags(appConfig *appConfig, logger *slog.Logger) error {
	command := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "migrations-directory",
				Value:       "migrations",
				Usage:       "Directory where SQL migrations can be found",
				Destination: &appConfig.MigrationsDirectory,
			},
			&cli.StringFlag{
				Name:        "db-connection-string",
				Value:       "postgres://init-microservice-go:password@localhost:5432/service_db?sslmode=disable",
				Usage:       "Connection String for connecting to a Postgres Database",
				Destination: &appConfig.DbConnectionString,
			},
		},
		Action: func(_ context.Context, _ *cli.Command) error {
			logger.Info("Config loaded.",
				slog.String("migrations-directory", appConfig.MigrationsDirectory),
				slog.String("db-connection-string", appConfig.DbConnectionString),
			)

			return nil
		},
	}

	err := command.Run(context.Background(), os.Args)

	// If no error, then it will return nil, which is fine.
	return fmt.Errorf("config failed to load from command flags: %w", err)
}
