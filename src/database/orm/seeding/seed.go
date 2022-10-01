package seeding

import (
	"log"

	"github.com/adiet95/gorent-api/src/database/orm"
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "start seeding",
	RunE:  seeder,
}

func seeder(cmd *cobra.Command, args []string) error {
	db, err := orm.New()
	if err != nil {
		return err
	}
	admin, _ := helpers.HashPassword("admin")
	user, _ := helpers.HashPassword("user")

	var datas = []models.User{
		{Email: "admin", Password: admin, Role: "admin"},
		{Email: "user", Password: user, Role: "user"},
	}

	if res := db.Create(&datas); res.Error != nil {
		return res.Error
	}
	log.Println("Seeding successful")
	return nil
}
