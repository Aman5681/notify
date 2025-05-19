package db

import (
	"fmt"
	"log"

	"github.com/Aman5681/notify/internal/config"
	"github.com/Aman5681/notify/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *config.Config) error {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmod=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = autoMigrateModels()

	if err != nil {
		return err
	}
	return nil

}

func autoMigrateModels() error {
	log.Println("📦 Running auto-migration...")

	modelsToMigrate := []interface{}{
		&models.User{}, // 👈 Add new models here
	}

	for _, model := range modelsToMigrate {
		if err := DB.AutoMigrate(model); err != nil {
			return err
		}
	}

	log.Println("✅ Auto-migration complete")
	return nil
}
