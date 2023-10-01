// controller/item_controller.go

package controller

import (
	"encoding/json"
	"fmt"
	"my-gorm-crud-app/model"
	"my-gorm-crud-app/view"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type ItemController struct {
	DB *gorm.DB
}

func NewItemController(db *gorm.DB) *ItemController {
	return &ItemController{DB: db}
}

func (c *ItemController) GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []model.Item
	c.DB.Find(&items)
	view.RespondJSON(w, http.StatusOK, items)
}

func (c *ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	var item model.Item
	if err := c.DB.First(&item, itemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			view.RespondJSON(w, http.StatusNotFound, "Item not found")
		} else {
			view.RespondJSON(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	view.RespondJSON(w, http.StatusOK, item)
}

func (c *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		view.RespondJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	c.DB.Create(&item)
	view.RespondJSON(w, http.StatusCreated, item)
}

func (c *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	var item model.Item
	if err := c.DB.First(&item, itemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			view.RespondJSON(w, http.StatusNotFound, "Item not found")
		} else {
			view.RespondJSON(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	var updatedItem model.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
		view.RespondJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	item.Name = updatedItem.Name
	item.Price = updatedItem.Price
	c.DB.Save(&item)
	view.RespondJSON(w, http.StatusOK, item)
}

func (c *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	var item model.Item
	if err := c.DB.First(&item, itemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			view.RespondJSON(w, http.StatusNotFound, "Item not found")
		} else {
			view.RespondJSON(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	c.DB.Delete(&item)
	view.RespondJSON(w, http.StatusNoContent, fmt.Sprintf("Item ID %s has been deleted", itemID))
}
