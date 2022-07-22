package main

import (
	"api-master-golang-products/controllers"
	"api-master-golang-products/database"
	"api-master-golang-products/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var TIME_OUT = time.Second * 5

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "JWT IS VALID")
}

func main() {

	pctx := context.Background()
	ctx, cancel := context.WithTimeout(pctx, TIME_OUT)

	defer cancel()
	// Load config files from application.json using Viper
	utils.LoadAppConfig()

	database.Connect(utils.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes paths
	RegisterProductRouters(ctx, router)

	// Start Server
	log.Printf("Starting Server on port %s", utils.AppConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", utils.AppConfig.Port), router)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8089",
		WriteTimeout: TIME_OUT,
		ReadTimeout:  TIME_OUT,
	}

	log.Println("Listen at port :8000")
	log.Fatal(srv.ListenAndServe())

}

func RegisterProductRouters(ctx context.Context, router *mux.Router) {
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.Handle("/api/validjwt", utils.ValidateJWT(Home))
	router.HandleFunc("/api/getjwt", utils.GetJwt)
}
