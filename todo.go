package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Index      int           `bson:"index"`
	Place      int           `bson:"place"`
	Task       int           `bson:"task"`
	State      int           `bson:"state"`
	CreateTime time.Time     `bson:"create_time"`
	StartTime  time.Time     `bson:"start_time"`
	Duration   time.Duration `bson:"duration"`
}

func (t *Todo) IsCompleted() bool {
	now := time.Now()
	now.Add(t.Duration)
	if now.After(t.StartTime) {
		return true
	}

	return false
}
