package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
)

type AppConfig struct {
	MigrationsDirectory string
	DbConnectionString  string
}

func loadAppConfigFromCommandFlags(appConfig *AppConfig, logger *slog.Logger) error {
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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			logger.Info("Config loaded.",
				slog.String("migrations-directory", appConfig.MigrationsDirectory),
				slog.String("db-connection-string", appConfig.DbConnectionString),
			)
			return nil
		},
	}

	err := command.Run(context.Background(), os.Args)
	return err // If no error, then it will return nil, which is fine.
}
