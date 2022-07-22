package main

import (
	"api-crud/controllers"
	"api-crud/database"
	"api-crud/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "JWT IS VALID")
}

func main() {
	// Load config files from application.json using Viper
	utils.LoadAppConfig()

	database.Connect(utils.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes paths
	RegisterProductRouters(router)

	// Start Server
	log.Printf("Starting Server on port %s", utils.AppConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", utils.AppConfig.Port), router)

}

func RegisterProductRouters(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.Handle("/api/validjwt", utils.ValidateJWT(Home))
	router.HandleFunc("/api/getjwt", utils.GetJwt)
}
