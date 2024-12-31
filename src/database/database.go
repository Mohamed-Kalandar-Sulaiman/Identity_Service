package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection based on the environment variable
func ConnectDatabase() error {
    var err error
    dbType := os.Getenv("DB_TYPE") // Use an environment variable to determine the database type

    switch dbType {
    case "sqlite":
        DB, err = gorm.Open(sqlite.Open("identity_service.db"), &gorm.Config{}) // SQLite file name
        if err != nil {
            return err
        }
        log.Println("Connected to the SQLite database.")

    case "postgres":
        dsn := "host=" + os.Getenv("PG_HOST") + 
                " user=" + os.Getenv("PG_USER") + 
                " password=" + os.Getenv("PG_PASSWORD") + 
                " dbname=" + os.Getenv("PG_DBNAME") + 
                " port=" + os.Getenv("PG_PORT") + 
                " sslmode=disable"
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            return err
        }
        log.Println("Connected to the PostgreSQL database.")

    default:
        return fmt.Errorf("unsupported database type: %s", dbType)
    }

    return nil
}
