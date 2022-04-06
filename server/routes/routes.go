package routes

import (
	"golang-backend/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/products", controllers.SampleProduct).Methods("GET")
	router.HandleFunc("/order/cart/add", controllers.CartAdd).Methods("POST")
	router.HandleFunc("/order/cart/get", controllers.FullCart).Methods("GET")
	router.HandleFunc("/order/cart/delete/{id}", controllers.DeleteCart).Methods("DELETE")
	router.HandleFunc("/{name}", controllers.EachProductGet).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	return router
}
