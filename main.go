package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/manchinagarjuna/bettertomorrow/pkg/event"
	"github.com/manchinagarjuna/bettertomorrow/pkg/organization"
	"github.com/manchinagarjuna/bettertomorrow/pkg/post"
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

var dbCollections util.Collections

var debugEnabled bool

func apiHandler(w http.ResponseWriter, r *http.Request) {

	if debugEnabled {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
	}

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
		user.UserHandler(w, requestMessage, dbCollections, r.Method)
	case "posts":
		post.PostHandler(w, requestMessage, dbCollections, r.Method)
	case "organizations":
		organization.OrgHandler(w, requestMessage, dbCollections, r.Method)
	case "events":
		event.EventHandler(w, requestMessage, dbCollections, r.Method)
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
	debugEnabled = false

	database, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbCollections.User = database.Collection("user")
	dbCollections.Organization = database.Collection("organization")
	dbCollections.Event = database.Collection("event")
	dbCollections.Post = database.Collection("post")

	server := http.NewServeMux()

	//Serving static assets
	fs := http.FileServer(http.Dir("assets/static"))
	server.HandleFunc("/api/v1/", apiHandler)
	server.Handle("/", fs)

	log.Println("Serving on port " + port)
	err = http.ListenAndServe(":"+port, logger(server))

	if err != nil {
		log.Fatal(err.Error())
	}
}
