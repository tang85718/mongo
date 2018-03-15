package mongo

import "gopkg.in/mgo.v2"

const (
	DB_GLOBAL = "global" // catalog global
	C_PLAYER  = "player"
	C_ACTOR   = "actors"
	DB_ASYLUM = "asylum" // catalog asylum
	DB_WILDER = "wilderness" // catalog wilderness
	C_PENDING = "pending"
	C_TODO    = "todo" // catalog actor-token
)

type MongoDB struct {
	Conn *mgo.Session
}

func (self *MongoDB) Dial(url string) error {
	ms, err := mgo.Dial("")
	if err != nil {
		panic("连接")
	}

	ms.SetMode(mgo.Monotonic, true)
	self.Conn = ms
	return err
}

func (self *MongoDB) CheckHealth() bool {
	if self.Conn == nil {
		return false
	}

	return self.Conn.Ping() != nil
}
