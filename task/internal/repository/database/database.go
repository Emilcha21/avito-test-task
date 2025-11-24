package database

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "github.com/lib/pq"
)

func InitDB(config config.Storage) (*gorm.DB, error) {
	sdn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.Host, config.Port, config.DbName, config.Username, config.Password, config.Sslmode)
	db, err := gorm.Open(postgres.Open(sdn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Cannot connect to database", err)
		return nil, err
	}

	err = migrateDB(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, nil
}

func migrateDB(db *gorm.DB) error {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(
		&models.Team{},
		&models.User{},
		&models.PullRequest{},
		&models.PRUsers{},
	)
	if err != nil {
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
