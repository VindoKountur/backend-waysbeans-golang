package database

import (
	"backendwaysbeans/models"
	"backendwaysbeans/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Profile{},
		&models.Addresses{},
		&models.Transaction{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration completed")
}
