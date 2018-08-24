package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 取得した各行を格納する
type Row struct {
	textColumn string
}

func main() {

	// DBへ接続
	db, _ := sql.Open("mysql", "testuser:password@tcp(localhost:3306)/test_db")
	defer db.Close()

	// Select文発行
	rows, _ := db.Query(`
      SELECT
            text_column
      FROM
			test_tbl
	  LIMIT
	  		10
	`)
	defer rows.Close()

	// 1行ずつ取得
	for rows.Next() {
		var r Row
		rows.Scan(&(r.textColumn))
		fmt.Println(r)
	}
}
