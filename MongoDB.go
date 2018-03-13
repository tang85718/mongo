package mongo

import "gopkg.in/mgo.v2"

const (
	DB_ROOT  = "root"
	C_PLAYER = "player"
	C_ACTOR  = "actors"
)

type MongoDB struct {
	ms *mgo.Session
}

func (self *MongoDB) Dial(url string) error {
	ms, err := mgo.Dial("")
	if err != nil {
		panic("连接")
	}

	ms.SetMode(mgo.Monotonic, true)
	self.ms = ms
	return err
}

func (self *MongoDB) CheckHealth() bool {
	if self.ms == nil {
		return false
	}

	return self.ms.Ping() != nil
}
