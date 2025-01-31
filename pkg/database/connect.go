package database

import (
	"fmt"

	"github.com/renatorrocha/minipay/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	// Verificar se as variáveis de ambiente estão definidas
	dbHost := config.Config("DATABASE_HOST")
	dbPort := config.Config("DATABASE_PORT")
	dbUser := config.Config("DATABASE_USER")
	dbPass := config.Config("DATABASE_PASSWORD")
	dbName := config.Config("DATABASE_NAME")

	// Verificar se alguma variável está vazia
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		panic("The environment variables are not set")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Falha ao conectar ao banco de dados: %v", err))
	}

	fmt.Println("Database connected successfully")
	DB.AutoMigrate()
	fmt.Println("Database migrated successfully")
}
