package user

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

func UserHandler(w http.ResponseWriter, msg util.RequestMessage, collections util.Collections, httpMethod string) {
	var data interface{}
	var err error
	errorString := ""

	switch msg.Operation {
	case "get":
		data, err = GetUsers(collections.User)
	case "new":
		var user User
		err = json.Unmarshal(msg.Data, &user)
		if err != nil {
			break
		}
		err = InsertUser(collections.User, user)
	case "delete":
		var user User
		err = json.Unmarshal(msg.Data, &user)
		if err != nil {
			break
		}
		err = DeleteUser(collections.User, user)
	}

	if err != nil {
		errorString = err.Error()
	}

	response := util.ResponseMessage{Data: data, Error: errorString}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

func GetUsers(collection *mongo.Collection) ([]User, error) {
	var users []User
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return users, err
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var user *User = &User{}
		err := cur.Decode(user)
		if err != nil {
			return users, err
		}
		users = append(users, *user)
	}
	err = cur.Err()
	return users, err
}

func InsertUser(collection *mongo.Collection, user User) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	userBson, err := bson.Marshal(user)
	if err != nil {
		return err
	}

	res, err := collection.InsertOne(ctx, userBson)
	id := res.InsertedID
	fmt.Println(id)

	return err
}

func DeleteUser(collection *mongo.Collection, user User) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	userBson, err := bson.Marshal(user)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, userBson)

	return err
}
