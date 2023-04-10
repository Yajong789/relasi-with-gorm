package controllers

import (
	"encoding/json"
	"net/http"
	"yajong/models"

	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "tidak bisa di decode", http.StatusBadRequest)
		return
	}

	result := db.Create(&category)
	if result.Error != nil {
		http.Error(w, "data tidak bisa ditambahkan", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var category []models.Category
	if err = db.Find(&category).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(category)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	w.Write(response)

}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	var category models.Category
	result := db.Preload("Product").First(&category, vars["id"])
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	var category models.Category
	result := db.First(&category, vars["id"])
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "gagal mendecode json", http.StatusBadRequest)
		return
	}

	result = db.Save(&category)
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	var category models.Category
	result := db.Delete(&category, vars["id"])
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "tidak bisa di decode", http.StatusBadRequest)
		return
	}

	result := db.Create(&product)
	if result.Error != nil {
		http.Error(w, "data tidak bisa ditambahkan", http.StatusBadRequest)
		return
	}

	result = db.Preload("Category").First(&product)
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(product)

}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var product []models.Product
	if err = db.Preload("Category").Find(&product).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	w.Write(response)
}

func GetProductByCategoryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var product []models.Product
	if err = db.Preload("Category").Where("category_id", id).Find(&product).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	w.Write(response)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	var product models.Product
	result := db.Preload("Category").First(&product, vars["id"])
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "gagal mendecode json", http.StatusBadRequest)
		return
	}

	if db.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		http.Error(w, "data tidak bisa diupdate", http.StatusBadRequest)
		return
	}

	result := db.Preload("Category").First(&product, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "tidak menemukan category", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(product)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.DB()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product

	if db.Delete(&product, "id = ?", id).RowsAffected == 0 {
		http.Error(w, "data tidak bisa dihapus", http.StatusBadRequest)
		return
	}

	w.Write([]byte("succes"))
}
