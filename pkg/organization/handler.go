package organization

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

func OrgHandler(w http.ResponseWriter, msg util.RequestMessage, collections util.Collections, httpMethod string) {
	var data interface{}
	var err error
	errorString := ""

	switch msg.Operation {
	case "get":
		data, err = GetOrgs(collections.Organization)
	case "insert":
		var org Organization
		err = json.Unmarshal(msg.Data, &org)
		if err != nil {
			break
		}
		err = InsertOrg(collections.Organization, org)
	case "delete":
		var org Organization
		err = json.Unmarshal(msg.Data, &org)
		if err != nil {
			break
		}
		err = DeleteOrg(collections.Organization, org)
	}

	if err != nil {
		errorString = err.Error()
	}

	response := util.ResponseMessage{Data: data, Error: errorString}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

func GetOrgs(collection *mongo.Collection) ([]Organization, error) {
	var orgs []Organization
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return orgs, err
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var org *Organization = &Organization{}
		err := cur.Decode(org)
		if err != nil {
			return orgs, err
		}
		orgs = append(orgs, *org)
	}
	err = cur.Err()
	return orgs, err
}

func InsertOrg(collection *mongo.Collection, org Organization) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	orgBson, err := bson.Marshal(org)
	if err != nil {
		return err
	}

	res, err := collection.InsertOne(ctx, orgBson)
	id := res.InsertedID
	fmt.Println(id)

	return err
}

func DeleteOrg(collection *mongo.Collection, org Organization) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	orgBson, err := bson.Marshal(org)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, orgBson)

	return err
}
