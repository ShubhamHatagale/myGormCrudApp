// routes/item_routes.go

package routes

import (
	"my-gorm-crud-app/controller"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func SetupItemRoutes(r *mux.Router, db *gorm.DB) {
	itemController := controller.NewItemController(db)

	r.HandleFunc("/items", itemController.GetAllItems).Methods("GET")
	r.HandleFunc("/items/{id}", itemController.GetItem).Methods("GET")
	r.HandleFunc("/items", itemController.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", itemController.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", itemController.DeleteItem).Methods("DELETE")
}
