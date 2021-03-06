package organization

import "encoding/json"

type Organization struct {
	ID             string          `json:"id" bson:"_id,omitempty"`
	Name           string          `json:"name" bson:"name"`
	Location       string          `json:"location" bson:"location"`
	Contacts       []string        `json:"contacts" bson:"contacts"`
	Volunteers     []string        `json:"volunteers" bson:"volunteers"`
	Icon           string          `json:"icon" bson:"icon"`
	Events         []string        `json:"events" bson:"events"`
	Approved       bool            `json:"approved" bson:"approved"`
	HiddenMetaData json.RawMessage `json:"metadata" bson:"metadata"`
}
