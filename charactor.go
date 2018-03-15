package mongo

import "gopkg.in/mgo.v2/bson"

type Charactor struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	PlayerToken string        `bson:"player_token"`
	Name        string        `bson:"name"`
	HP          int           `bson:"hp"`
	Energy      int           `bson:"energy"`
	EnergyType  int           `bson:"energy_type"`
}

func (c *Charactor) ToDB(mgo *MongoDB, sdb string) error {
	db := mgo.Conn.DB(sdb).C(C_ACTOR)
	return db.Insert(c)
}

/**
	调用前确保ID不为空, 该函数会将数据从数据库中拿出来，然后删除数据库中的数据.
*/
func (c *Charactor) RemoveByID(mgo *MongoDB, sdb string) error {
	db := mgo.Conn.DB(sdb).C(C_ACTOR)
	err := db.FindId(c.ID).One(c)
	if err != nil {
		return err
	}

	return db.RemoveId(c.ID)
}
