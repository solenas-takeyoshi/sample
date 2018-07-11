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
}

func (c Authenticate) Index() revel.Result {
	con, err := sql.Open("mysql", "root:root@/asm_db")
	if err != nil {
		revel.ERROR.Println("FATAL", err)
	}
	defer con.Close()

	// select
	rows, err := con.Query("select id, username from user_t")
	if err != nil {
		revel.ERROR.Println("FATAL", err)
	}
	items := make([]*UserData, 0)

	var id uint
	var username string
	for rows.Next() {
		err = rows.Scan(&id, &username)
		if err != nil {
			revel.ERROR.Println("FATAL", err)
		}

		items = append(items, &UserData{id, username})
	}

	// Insert
	// _, err = con.Exec("insert into USER_T (username) values (?)", "test")
	// if err != nil {
	// 	revel.ERROR.Println("FATAL", err)
	// }
	return c.Redirect("menu")
}
