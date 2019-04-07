package util

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)

type RequestMessage struct {
	Operation string          `json:operation`
	Data      json.RawMessage `json:data`
}

type ResponseMessage struct {
	Data  interface{} `json:data`
	Error string      `json:error`
}

type Collections struct {
	User         *mongo.Collection
	Organization *mongo.Collection
	Post         *mongo.Collection
	Event        *mongo.Collection
}
