package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	// MongoDBサーバへの接続
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	// DBオブジェクトの取得
	db := session.DB("test_db")

	p := new(struct {
		ID   bson.ObjectId `bson:"_id"`
		Name string        `bson:"name"`
		Age  int           `bson:"age"`
	})
	query := db.C("test_cl").Find(bson.M{})
	query.One(&p)

	/**
	 * 結果
	**/
	fmt.Printf("%+v\n", p)
	// &{ID:ObjectIdHex("5478517a9871b9b8e42e2ee2") Name:田井中律 Age:17}
}
