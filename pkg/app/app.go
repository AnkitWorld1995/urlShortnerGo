package app

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urlShortnerGo/config"
	"github.com/urlShortnerGo/pkg/domain"
	"github.com/urlShortnerGo/pkg/handler"
	"github.com/urlShortnerGo/pkg/service"
	"log"
	"net/http"
)

func StartApplication(cfg *config.AppConfig) {

	// Start an AWS Session
	newAWSSession, err := session.NewSession(&aws.Config{
		Region: &config.GetConfig().AwsConfigure.AwsRegion,
	})
	if err != nil {
		log.Fatalf("the AWS Client Failed with an Error %s", err.Error())
		return
	}
	log.Println("------- appConfig AWS-----------", config.GetConfig().AwsConfigure.AwsRegion, newAWSSession.Config.Region)

	// Start an AWS DynamoDB Session
	newDynamoDBSession := dynamodb.New(newAWSSession, aws.NewConfig().WithEndpoint(fmt.Sprintf("%s%s", cfg.AwsConfigure.Dynamodb.Address, cfg.AwsConfigure.Dynamodb.Port)))

	router := mux.NewRouter()
	pingHandler := handler.PingHandler{}
	urlHandler := handler.UrlHandler{Service: service.NewUrlService(domain.NewShortnerRepo(newDynamoDBSession))}

	router.HandleFunc("/ping", pingHandler.Greet).Methods(http.MethodGet)
	router.HandleFunc("/get-url", urlHandler.GetUrl).Methods(http.MethodGet)

	//Setup cors
	router.Use(mux.CORSMethodMiddleware(router))
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Origin", "Host", "Access-Control-Allow-Credentials", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	// starting server
	port := cfg.GeneralConfig.AppPort
	environment := cfg.GeneralConfig.Environment
	log.Println(fmt.Sprintf("Starting server on localhost:%s in: %s", port, environment))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
