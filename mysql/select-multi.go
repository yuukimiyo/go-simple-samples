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

	// Select文発行
	rows, _ := db.Query("SELECT id, text_column FROM test_tbl")
	defer rows.Close()

	// 1行ずつ取得
	for rows.Next() {
		id, textColumn := 0, ""
		rows.Scan(&id, &textColumn)
		fmt.Printf("id=%d, text_column='%s'\n", id, textColumn)
	}
}
