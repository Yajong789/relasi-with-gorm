package main

import (
	"fmt"
	"net/http"
	"yajong/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", controllers.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", controllers.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/products/bycategoryid/{id}", controllers.GetProductByCategoryId).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	fmt.Println("Starting server at localhost:8080")
	http.ListenAndServe(":8080", router)

}
