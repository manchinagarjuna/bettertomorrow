package event

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Org         string             `json:"org" bson:"org"`
	Starttime   string             `json:"startTime" bson:"startTime"`
	Endtime     string             `json:"endTime" bson:"endTime"`
	Location    string             `json:"location" bson:"location"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	Numlikes    int64              `json:"numLikes" bson:"numLikes"`
	Reports     []string           `json:"reports" bson:"reports"`
	Comments    []string           `json:"comments" bson:"comments"`
	Posttime    string             `json:"postTime" bson:"postTime"`
}

type GetEventsByOrganizationRequest struct {
	ID string `json:id`
}
