package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Row struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

func main() {

	// MongoDBサーバへの接続
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	// DBオブジェクトの取得
	db := session.DB("test_db")

	// Collectionオブジェクトの取得
	col := db.C("test_cl")

	// Insertする行を作成
	row := Row{
		ID:   bson.NewObjectId(),
		Name: "my name",
		Age:  20,
	}

	// Insertを実行
	err := col.Insert(&row)
	if err != nil {
		// エラー処理
		panic(err.Error())
	}
}
