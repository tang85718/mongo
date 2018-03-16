package mongo

import "gopkg.in/mgo.v2/bson"

const (
	PLACE_GOD_SPACE  = 0
	PLACE_ASYLUM     = 1
	PLACE_WILDERNESS = 2
)

/**
Place:
0: 刚出生的baby
1: 进入庇护所Asylum
2: 进入荒原
 */
type Charactor struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	PlayerToken string        `bson:"player_token"`
	Name        string        `bson:"name"`
	HP          int           `bson:"hp"`
	Energy      int           `bson:"energy"`
	EnergyType  int           `bson:"energy_type"`
	Place       int           `bson:"place"`
}

func (c *Charactor) ToDB(mgo *MongoDB, sdb string) error {
	db := mgo.Conn.DB(sdb).C(C_ACTOR)
	return db.Insert(c)
}

func (c *Charactor) UpdatePlace(mgo *MongoDB, value int) error {
	db := mgo.Conn.DB(DB_GLOBAL).C(C_ACTOR)
	m := bson.M{"$set": bson.M{"place": value}}
	return db.UpdateId(c.ID, m)
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
