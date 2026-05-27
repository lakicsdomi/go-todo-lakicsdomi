package routes

import (
	"go-todo/controllers"

	"github.com/gorilla/mux"
	"github.com/lakicsdomi/argus"
)

func Init() *mux.Router {
	logger, _ := argus.Init("logs")
	logger.Verbose.Log("ROUTER", "Initializing routes...")
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Show)
	router.HandleFunc("/add", controllers.Add).Methods("POST")
	router.HandleFunc("/delete/{id}", controllers.Delete)
	router.HandleFunc("/complete/{id}", controllers.Complete)

	logger.Verbose.Log("ROUTER", "Routes initialized.")
	return router
}
