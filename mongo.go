package connection

import (
	"os"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var mongo_session *mgo.Session

func Mongo() *mgo.Database {
	if mongo_session == nil {
		mongo_session = MongoSession()
	}

	return mongo_session.DB(os.Getenv("DB_MONGO_DATABASE"))
}
	
func MongoSession() *mgo.Session {
	mongo_addresses := strings.Split(os.Getenv("DB_MONGO_HOST"), ",")

	info := &mgo.DialInfo{
		Addrs:    mongo_addresses,
		Timeout:  30 * time.Second,
		Database: os.Getenv("DB_MONGO_DATABASE"),
		Username: os.Getenv("DB_MONGO_USERNAME"),
		Password: os.Getenv("DB_MONGO_PASSWORD"),
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	defer func() *mgo.Session {
		if r := recover(); r != nil {
			panic_session := MongoSession()
			panic_session.Refresh()

			return panic_session.Clone()
		}
		return nil
	}()

	return session
}
