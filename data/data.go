package data

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Item struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Products  []string      `bson:"products" json:"products"`
	Date      string        `bson:"date" json:"date"`
	Source    string        `bson:"source" json:"source"`
	TimeAdded time.Time     `bson:"time_added" json:"time_added"`
}
