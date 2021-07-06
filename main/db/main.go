package main

import (
	"fmt"
	"go_setup_v1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_setup_v1 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DropTables(db)
	CreateCategoriesTable(db)
	CreateUsersTable(db)
}

func DropTables(db *gorm.DB) {
	//db.DropTableIfExists(products.Product{})
	db.Migrator().DropTable(models.Category{})
	fmt.Println("Tables dropped")
}

func CreateCategoriesTable(db *gorm.DB) {
	db.AutoMigrate(&models.Category{})
	fmt.Println("Categories table created")
}

func CreateUsersTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	fmt.Println("Users table created")
}

/*
func CreateProductsTable(db *gorm.DB) {
	db.AutoMigrate(&products.Product{})
	db.Model(&products.Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")

	fmt.Println("Products table created")
}
*/
