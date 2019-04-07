package post

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Org      string             `json:"orgId" bson:"orgId"`
	Title    string             `json:title bson:"title"`
	Text     string             `json:text bson:"text"`
	Image    string             `json:image bson:"image"`
	IsEvent  bool               `json:isEvent bson:"isEvent"`
	NumLikes uint64             `json:numLikes bson:"numLikes"`
}

type GetPostsByOrganizationRequest struct {
	ID string `json:id`
}
