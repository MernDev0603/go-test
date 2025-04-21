package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB
var err error

// Initialize database connection
func Init() {
	// Set up MySQL connection
	// Replace with your actual MySQL connection details
	// dsn := "root:@tcp(127.0.0.1:3306)/bowatt?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("root:@tcp(%s:3306)/bowatt?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_HOST"))
	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to the database!")
	}

	// Automatically migrate the schema
	fmt.Println("Database connected successfully!")
}
