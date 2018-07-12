package controllers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

type Authenticate struct {
	*revel.Controller
}

type UserData struct {
	id       uint
	username string
	address  string
}

func (c Authenticate) Index() revel.Result {
	con, err := sql.Open("mysql", "root:root@/asm_db")
	if err != nil {
		revel.ERROR.Println("FATAL", err)
	}
	defer con.Close()

	// ----------------------------------------------------------------------------------
	// // select1
	// // パスワード取得
	// var password string
	// userID := "take"
	// err = con.QueryRow("select password from user_t where userid = ?", userID).Scan(&password)
	// if err != nil {
	// 	revel.ERROR.Println("FATAL", err)
	// 	panic(err.Error)
	// }

	// ----------------------------------------------------------------------------------
	// select2
	// 複数取得( + 複合化)
	pass := "address-Pass"
	rows, err := con.Query("select id, username, convert( AES_DECRYPT(UNHEX(address), ?) USING utf8 ) as address from user_t", pass)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err.Error)
	}
	items := make([]*UserData, 0)

	var id uint
	var username string
	var address string
	for rows.Next() {
		err = rows.Scan(&id, &username, &address)
		if err != nil {
			revel.ERROR.Println("FATAL", err)
			panic(err.Error)
		}

		items = append(items, &UserData{id, username, address})
	}

	// ----------------------------------------------------------------------------------
	// // Insert
	// stmtIns, err := con.Prepare("insert into USER_T (userid, username, address, password) values (?, ?,hex(aes_encrypt(?,?)), sha2(?,?))")
	// if err != nil {
	// 	revel.ERROR.Println("FATAL", err)
	// 	panic(err.Error)
	// }
	// // TODO:日本語が登録できない。
	// // address := "〒100-8111 東京都千代田区千代田１−１"
	// userid := "take"
	// username := "takeyoshi"
	// address := "test"
	// key := "address-Pass"
	// password := "password"
	// hashlength := 256
	// _, err = stmtIns.Exec(userid, username, address, key, password, hashlength)
	// if err != nil {
	// 	revel.ERROR.Println("FATAL", err)
	// 	panic(err.Error)
	// }

	return c.Redirect("menu")
}
