package linksModel

import (
	"gopkg.in/mgo.v2/bson"

)

type Links struct {
	ObjId 	bson.ObjectId		`bson:"_id"`
	Id 		uint32
	LinkName	string	//链接
	Groupings	bson.ObjectId	//分组
}