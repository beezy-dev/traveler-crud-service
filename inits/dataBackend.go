package inits

import (
	"fmt"
	"log"
	"os"
	"slices"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DataBackEnd() {
	// loading variables from .env file
	LoadEnv()

	// defining useful vars
	var err error
	dbType := os.Getenv("DB_TYPE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// supported backend
	backends := []string{"postgres", "mariadb", "sqlite"}
	if !slices.Contains(backends, dbType) {
		log.Fatal("Fatal: ", dbType, " is not support. Refer to documentation for supported databackend")
	}

	// dsn connection based on backend type to initiate backend db
	switch dbType {
	case "postgres":
		log.Println("Info: database backend service configured with:", dbType)
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s", dbHost, dbPort, dbUser, dbPass)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println("Error: can not connect to", dbType, "service.")
			log.Fatal("Fatal: ", err)
		} else {
			log.Println("Info: connected to", dbType, "service.")
		}
		CreateDBPostgres()
	case "mariadb":
		log.Println("Info: database backend service configured with:", dbType)
		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 	Logger: logger.Default.LogMode(logger.Silent),
		// })
		log.Fatal("Fatal: database backend service configured with: ", dbType, " not yet implemented")
	}

	// dsn connection based on backend type to connect to db
	switch dbType {
	case "postgres":
		log.Println("Info: reconnecting to service", dbType, "and opening", dbName)
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPass, dbName)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println("Error: can not connect to database", dbName, "in service", dbType)
			log.Fatal(err)
		} else {
			log.Println("Info: connected to database", dbName, "in service", dbType)
			// updating backend schema for the given database - TODO verify the impact if structure change
			log.Println("Info: updating schema of database", dbName, "in service", dbType)
			MigrateSchema()
		}
	case "mariadb":
		log.Println("Info: database backend service configured with:", dbType)
		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 	Logger: logger.Default.LogMode(logger.Silent),
		// })
		log.Fatal("Fatal: database backend service configured with: ", dbType, " not yet implemented")
	case "sqlite":
		log.Println("Info: database backend service configured with:", dbType)
		DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		log.Fatal("Fatal: database backend service configured with: ", dbType, " not yet implemented")
	}

}
