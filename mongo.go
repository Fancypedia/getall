package peda

import (
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: MONGOCONNSTRINGENV,
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetAllBangunanLineString(mongoconn *mongo.Database, collection string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mongoconn, collection)
	return lokasi
}

func GetAllUser(mongoconn *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](mongoconn, collection)
	return user
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}
