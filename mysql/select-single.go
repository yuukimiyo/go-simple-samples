package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// DBへ接続
	db, _ := sql.Open("mysql", "testuser:password@tcp(localhost:3306)/test_db")
	defer db.Close()

	// Select文を発行
	result := db.QueryRow("SELECT id, text_column FROM test_tbl WHERE id = 1")

	// 結果を取得して表示
	id, textColumn := 0, ""
	result.Scan(&id, &textColumn)
	fmt.Printf("id=%d, text_column='%s'\n", id, textColumn)
}
