package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

/**
Todo.Task
0: 游戏刚开始

Todo.Place
0: asylum
1: wilderness
 */

type Todo struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	ActorToken string        `bson:"actor_token"`
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

func (t *Todo) ToDB(mgo *MongoDB) error {
	db := mgo.Conn.DB(DB_GLOBAL).C(C_TODO)
	count, err := db.Find(bson.M{"actor_token": t.ActorToken}).Count()
	if err != nil {
		return err
	}
	t.Index = count + 1
	return db.Insert(t)
}

func GetAllToDo(mgo *MongoDB) ([]Todo, error) {
	db := mgo.Conn.DB(DB_GLOBAL).C(C_TODO)

	results := []Todo{}
	if err := db.Find(nil).All(&results); err != nil {
		return nil, err
	}

	return results, nil
}
