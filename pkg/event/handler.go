package event

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/manchinagarjuna/bettertomorrow/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func EventHandler(w http.ResponseWriter, msg util.RequestMessage, collections util.Collections, httpMethod string) {
	var data interface{}
	var err error
	errorString := ""

	switch msg.Operation {
	case "get":
		data, err = GetEvents(collections.Event, "")
	case "new":
		var event Event
		err = json.Unmarshal(msg.Data, &event)
		if err != nil {
			break
		}
		err = InsertEvent(collections.Event, event)
	case "delete":
		var event Event
		err = json.Unmarshal(msg.Data, &event)
		if err != nil {
			break
		}
		err = DeleteEvent(collections.Event, event)
	case "filterByOrg":
		var getEventsReq GetEventsByOrganizationRequest
		err = json.Unmarshal(msg.Data, &getEventsReq)
		if err != nil {
			break
		}
		data, err = GetEvents(collections.Post, "{\"orgId\""+getEventsReq.ID+"}")
	}

	if err != nil {
		errorString = err.Error()
	}

	response := util.ResponseMessage{Data: data, Error: errorString}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

func GetEvents(collection *mongo.Collection, filter string) ([]Event, error) {
	var events []Event
	var bdoc interface{}
	err := bson.UnmarshalExtJSON([]byte(filter), false, &bdoc)
	if err != nil {
		panic(err)
	}

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return events, err
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var event *Event = &Event{}
		err := cur.Decode(event)
		if err != nil {
			return events, err
		}
		events = append(events, *event)
	}
	err = cur.Err()
	return events, err
}

func InsertEvent(collection *mongo.Collection, event Event) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	eventBson, err := bson.Marshal(event)
	if err != nil {
		return err
	}

	res, err := collection.InsertOne(ctx, eventBson)
	id := res.InsertedID
	fmt.Println(id)

	return err
}

func DeleteEvent(collection *mongo.Collection, event Event) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	eventBson, err := bson.Marshal(event)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, eventBson)

	return err
}
