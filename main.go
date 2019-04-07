package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/manchinagarjuna/bettertomorrow/pkg/user"
	"github.com/manchinagarjuna/bettertomorrow/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Basic logger to help with debugging missing static page errors.
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

var userCollection *mongo.Collection
var orgCollection *mongo.Collection
var eventCollection *mongo.Collection

func apiHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var requestMessage util.RequestMessage
	err := decoder.Decode(&requestMessage)
	if err != nil {
		response := util.ResponseMessage{Data: "", Error: err.Error()}
		res, _ := json.Marshal(response)
		w.Write([]byte(res))
		return
	}

	switch strings.TrimRight(strings.TrimPrefix(r.URL.Path, "/api/v1/"), "/") {
	case "users":
		go user.UserHandler(w, requestMessage, userCollection, r.Method)
	}
}

func connectToMongoDB() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://mongouser:bettertomorrow@bettertomorrow-ffmrf.gcp.mongodb.net/test?retryWrites=true"))

	if err != nil {
		return nil, err
	}

	database := client.Database("bettertomorrow")

	return database, nil
}

func main() {

	port := "8080"

	database, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	userCollection = database.Collection("user")
	orgCollection = database.Collection("organization")
	eventCollection = database.Collection("event")

	// Creating a ServeMux to log messages as well as serve static pages.
	server := http.NewServeMux()

	//Serving static assets
	fs := http.FileServer(http.Dir("assets/static"))
	// Connect a basic logger.
	handlerWithLogger := logger(fs)
	server.Handle("/", fs)
	http.HandleFunc("/api/v1/", apiHandler)

	log.Println("Serving on port " + port)
	err = http.ListenAndServe(":"+port, handlerWithLogger)

	if err != nil {
		log.Fatal(err.Error())
	}
}
