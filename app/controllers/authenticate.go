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

// func (c Authenticate) Index(inputUserID string, inputPassword string) revel.Result {
func (c Authenticate) Index() revel.Result {
	inputUserID := c.Params.Get("inputUserID")
	inputPassword := c.Params.Get("inputPassword")
	remembermeV := c.Params.Get("remembermeV")
	remembermeN := c.Params.Get("remembermeN")
	remembermeI := c.Params.Get("remembermeI")
	buttonsubmit := c.Params.Get("button_submit")

	revel.WARN.Println("inputUserID:" + inputUserID)
	revel.WARN.Println("inputPassword:" + inputPassword)
	revel.WARN.Println("remembermeV:" + remembermeV)
	revel.WARN.Println("remembermeN:" + remembermeN)
	revel.WARN.Println("remembermeI:" + remembermeI)
	revel.WARN.Println("buttonsubmit:" + buttonsubmit)

	// コントロール受け取り
	//hidden
	testHiddenName := c.Params.Get("testHiddenName")
	revel.WARN.Println("testHiddenName:" + testHiddenName)
	//テキストボックス
	testTextName := c.Params.Get("testTextName")
	revel.WARN.Println("testTextName:" + testTextName)
	//チェックボックス
	testCheckboxName := c.Params.Get("testCheckboxName")
	revel.WARN.Println("testCheckboxName:" + testCheckboxName)
	//チェックボックス(複数)
	len := len(c.Params.Form["testCheckboxName2"])
	for i := 0; i < len; i++ {
		revel.WARN.Println("testCheckboxName2:" + c.Request.Form["testCheckboxName2"][i])
	}

	//ラジオボタン
	testRadioName := c.Params.Get("testRadioName")
	revel.WARN.Println("testRadioName:" + testRadioName)
	//ラジオボタン(複数)
	testRadioName2 := c.Params.Get("testRadioName2")
	revel.WARN.Println("testRadioName2:" + testRadioName2)
	//ファイル
	testFileName := c.Params.Get("testFileName")
	revel.WARN.Println("testFileName:" + testFileName)
	//数値(HTML5から追加)
	testNumberName := c.Params.Get("testNumberName")
	revel.WARN.Println("testNumberName:" + testNumberName)
	//サブミット
	testSubmitName := c.Params.Get("testSubmitName")
	revel.WARN.Println("testSubmitName:" + testSubmitName)
	//画像ファイル
	testImageName := c.Params.Get("testImageName")
	revel.WARN.Println("testImageName:" + testImageName)
	//電話番号(HTML5から追加)
	testTelName := c.Params.Get("testTelName")
	revel.WARN.Println("testTelName:" + testTelName)
	//email(HTML5から追加)
	testEmailName := c.Params.Get("testEmailName")
	revel.WARN.Println("testEmailName:" + testEmailName)
	//レンジ(HTML5から追加)
	testRrangeName := c.Params.Get("testRrangeName")
	revel.WARN.Println("testRrangeName:" + testRrangeName)
	//検索テキスト(HTML5から追加)
	testSearchName := c.Params.Get("testSearchName")
	revel.WARN.Println("testSearchName:" + testSearchName)
	//url(HTML5から追加)
	testURLName := c.Params.Get("testUrlNameName")
	revel.WARN.Println("testUrlName:" + testURLName)
	//UTC（協定世界時）(HTML5から追加)
	testDatetimeName := c.Params.Get("testDatetimeName")
	revel.WARN.Println("testDatetimeName:" + testDatetimeName)
	//日付(HTML5から追加)
	testDateName := c.Params.Get("testDateName")
	revel.WARN.Println("testDateName:" + testDateName)
	//月(HTML5から追加)
	testMonthName := c.Params.Get("testMonthName")
	revel.WARN.Println("testMonthName:" + testMonthName)
	//週(HTML5から追加)
	testWeekName := c.Params.Get("testWeekName")
	revel.WARN.Println("testWeekName:" + testWeekName)
	//時間(HTML5から追加)
	testTimeName := c.Params.Get("testTimeName")
	revel.WARN.Println("testTimeName:" + testTimeName)
	//ローカル日時(HTML5から追加)
	testDatetimelocalName := c.Params.Get("testDatetimelocalName")
	revel.WARN.Println("testDatetimelocalName:" + testDatetimelocalName)
	//色(HTML5から追加)
	testColorName := c.Params.Get("testColorName")
	revel.WARN.Println("testColorName:" + testColorName)
	//汎用ボタン
	testButtonName := c.Params.Get("testButtonName")
	revel.WARN.Println("testButtonName:" + testButtonName)

	// ----------------------------------------------------------------------------------
	// ----------------------------------------------------------------------------------
	// DBアクセス
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
	// // select2
	// // 複数取得( + 複合化)
	// pass := "address-Pass"
	// rows, err := con.Query("select id, username, convert( AES_DECRYPT(UNHEX(address), ?) USING utf8 ) as address from user_t", pass)
	// if err != nil {
	// 	revel.ERROR.Println("FATAL", err)
	// 	panic(err.Error)
	// }
	// items := make([]*UserData, 0)

	// var id uint
	// var username string
	// var address string
	// for rows.Next() {
	// 	err = rows.Scan(&id, &username, &address)
	// 	if err != nil {
	// 		revel.ERROR.Println("FATAL", err)
	// 		panic(err.Error)
	// 	}

	// 	items = append(items, &UserData{id, username, address})
	// }

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
