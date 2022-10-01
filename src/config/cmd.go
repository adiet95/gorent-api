package config

import (
	"github.com/adiet95/gorent-api/src/database/orm"
	"github.com/adiet95/gorent-api/src/database/orm/seeding"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "Simple Backend Gorent With golang",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
	initCommand.AddCommand(seeding.SeedCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
