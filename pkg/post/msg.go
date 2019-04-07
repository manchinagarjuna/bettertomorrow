package post

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	org      string             `json:"orgId" bson:"orgId"`
	Title    string             `json:title bson:"title"`
	Data     string             `json:data bson:"data"`
	IsEvent  bool               `json:isEvent bson:"isEvent"`
	NumLikes uint64             `json:numLikes bson:"numLikes"`
}
