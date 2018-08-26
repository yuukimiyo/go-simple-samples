package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// DBへ接続
	// [username[:password]@][tcp(host[:port])]/dbname[param]
	db, _ := sql.Open("mysql", "testuser:password@tcp(localhost:3306)/test_db?collation=utf8mb4_bin")
	defer db.Close()

	// Select文発行
	row := db.QueryRow("SELECT id, name FROM test_tbl WHERE id = 1")

	// 結果取得用の変数を初期化
	id, name := 0, ""

	// 結果を取得して表示
	row.Scan(&id, &name)
	fmt.Printf("id=%d, name='%s'\n", id, name)
}
