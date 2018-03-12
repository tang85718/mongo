package mongo

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB_ROOT  = "root"
	C_PLAYER = "player"
	C_ACTOR  = "actors"
)

type Player struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	DisplayID   string        `bson:"id"`
	Token       string        `bson:"token"`
	Phone       string        `bson:"phone"`
	Country     string        `bson:"country"`
	Province    string        `bson:"province"`
	City        string        `bson:"city"`
	Lat         float64       `bson:"lat"`
	Lng         float64       `bson:"lng"`
	Birthday    time.Time     `bson:"birthday"`
	CreateTime  time.Time     `bson:"create_time"`
	UpdateTime  time.Time     `bson:"update_time"`
	Platform    int           `bson:"platform"` // 1:iOS 2:Android
	FCM         bool          `bson:"fcm"` // 如果是Android机器，是否支持FCM服务，iOS忽略
	DeviceToken string        `bson:"dev_token"`
}
