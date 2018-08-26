package main

import (
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

	// Collectionオブジェクトの取得
	col := db.C("test_cl")

	// Insert用のJson文字列（のByte配列）を作成
	j := []byte(`{
		"name": "my name",
		"age": 20
	}`)

	// 上記Byte配列をBsonへUnmarshal
	var b interface{}
	bson.UnmarshalJSON(j, &b)

	// Insertを実行
	err := col.Insert(&b)
	if err != nil {
		// エラー処理
		panic(err.Error())
	}
}
