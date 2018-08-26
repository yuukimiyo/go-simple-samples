package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// DBへ接続
	// [username[:password]@][tcp(host[:port])]/dbname[param]
	db, _ := sql.Open("mysql", "testuser:password@tcp(localhost:3306)/test_db?interpolateParams=true&collation=utf8mb4_bin")
	defer db.Close()

	// SQL文を発行
	_, err := db.Exec("insert into test_tbl (name) values (?)", "my name")
	if err != nil {
		// エラー処理
		panic(err.Error())
	}
}
