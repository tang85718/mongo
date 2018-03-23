package mongo

import "gopkg.in/mgo.v2/bson"

type Map struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
