package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//test:test => user:password
	//db:3306 => alias:port
	///test => database name
	db, err := sql.Open("mysql", "test:test@tcp(db:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//1レコード取得
	// var email string
	// var name string
	// if err := db.QueryRow("SELECT email, name FROM users WHERE id = ?", 2).Scan(&email, &name); err != nil {
	// 	fmt.Fprintf(w, err.Error())
	// 	return
	// }
	// result := name + ", " + email
	// fmt.Fprintf(w, result)

	//SELECT * FROM users WHERE id = ? LIMIT 1;

	//存在チェック
	var result int
	if err := db.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE id = ?)", 2).Scan(&result); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//create(insert)
	rlt, err := db.Exec("INSERT INTO users(user_id, name, email) VALUES(?, ?, ?)", "alskdfj987", "tarou tanaka", "taroutest@mail.com")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//レコード登録後のIDの取得
	lastid, err := rlt.LastInsertId()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//RowsAffected : クエリが影響を与えた行数を返す
	rowCnt, err := rlt.RowsAffected()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "result : %d, lastID : %d, affected : %d", result, lastid, rowCnt)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
