package orm

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// create table
			ID: "0001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}, &models.Histories{}, &models.Vehicle{})
			},
			// drop table
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.User{}, &models.Histories{}, &models.Vehicle{})
			},
		},
	})
	if migUp {
		if err := m.Migrate(); err != nil {
			return err
		}
		log.Println("Migration up done")
		return nil
	}
	if migDown {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		log.Println("Rollback up done")
		return nil
	}
	log.Println("init schema database done")
	return nil
}
