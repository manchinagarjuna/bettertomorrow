package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        string             `json:"userId" bson:"userId"`
	Displayname   string             `json:displayName bson:"displayName"`
	Icon          string             `json:icon bson:"icon"`
	Location      string             `json:location bson:"location"`
	RadiusInMiles uint16             `json:radiusInMiles bson:"radiusInMiles"`
	Orgs          []string           `json:orgs bson:"orgs"`
}
