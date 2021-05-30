package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"wedding/ent"
)

var GenerateMigrations = cli.Command{
	Name:      "gen-migrations",
	ArgsUsage: "[> ./output.sql]",
	Action: func(ctx *cli.Context) error {
		config, err := ProvideEntConfig()
		if err != nil {
			return err
		}

		// Connect to db.
		client, err := ent.Open("postgres", config.ConnectionString, ent.Log(log.Println))
		if err != nil {
			return err
		}
		defer client.Close()

		// Dump migration changes to stdout.
		return client.Schema.WriteTo(ctx.Context, os.Stdout)
	},
}
