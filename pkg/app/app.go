package app

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urlShortnerGo/config"
	"github.com/urlShortnerGo/pkg/handler"
	"log"
	"net/http"
)

func StartApplication(cfg *config.AppConfig) {

	router := mux.NewRouter()
	pingHandler := handler.PingHandler{}

	router.HandleFunc("/ping", pingHandler.Greet).Methods(http.MethodGet)

	//Setup cors
	router.Use(mux.CORSMethodMiddleware(router))
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Origin", "Host", "Access-Control-Allow-Credentials", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	// starting server
	port := cfg.Config.AppPort
	environment := cfg.Config.Environment
	log.Println(fmt.Sprintf("Starting server on localhost:%s in: %s", port, environment))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
