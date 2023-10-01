package main

import (
	"log"
	"net/http"

	"my-gorm-crud-app/model"
	"my-gorm-crud-app/routes"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func main() {
	// Initialize the database connection
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the model
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.Users{})

	r := mux.NewRouter()

	// Set up routes for the Item resource
	routes.SetupItemRoutes(r, db)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
