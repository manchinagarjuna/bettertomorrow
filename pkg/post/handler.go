package post

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

func PostHandler(w http.ResponseWriter, msg util.RequestMessage, collections util.Collections, httpMethod string) {
	var data interface{}
	var err error
	errorString := ""

	switch msg.Operation {
	case "get":
		data, err = GetPosts(collections.Post, "")
	case "new":
		var post Post
		err = json.Unmarshal(msg.Data, &post)
		if err != nil {
			break
		}
		err = InsertPost(collections.Post, post)
	case "delete":
		var post Post
		err = json.Unmarshal(msg.Data, &post)
		if err != nil {
			break
		}
		err = DeletePost(collections.Post, post)
	case "filterByOrg":
		var getPostsReq GetPostsByOrganizationRequest
		err = json.Unmarshal(msg.Data, &getPostsReq)
		if err != nil {
			break
		}
		data, err = GetPosts(collections.Post, "{\"orgId\""+getPostsReq.ID+"}")
	}

	if err != nil {
		errorString = err.Error()
	}

	response := util.ResponseMessage{Data: data, Error: errorString}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

func GetPosts(collection *mongo.Collection, filter string) ([]Post, error) {
	var posts []Post

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return posts, err
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var post *Post = &Post{}
		err := cur.Decode(post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, *post)
	}
	err = cur.Err()
	return posts, err
}

func InsertPost(collection *mongo.Collection, post Post) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	postBson, err := bson.Marshal(post)
	if err != nil {
		return err
	}

	res, err := collection.InsertOne(ctx, postBson)
	id := res.InsertedID
	fmt.Println(id)

	return err
}

func DeletePost(collection *mongo.Collection, post Post) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	postBson, err := bson.Marshal(post)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, postBson)

	return err
}
