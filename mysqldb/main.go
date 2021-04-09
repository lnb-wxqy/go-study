package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

func main() {
	userName := "root"
	password := "Wxqy@213"
	dsn := "39.106.0.106:3306"
	initCollectDB(userName, password, dsn)
	err := runCache()

	fmt.Println("er: ",err)

	os.Exit(0)
}

var Conn *sqlx.DB

func initCollectDB(userName, password, dsn string) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/interview?multiStatements=true&charset=utf8", userName, password, dsn)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		fmt.Println("连接数据库失败： ", err)
		return
	}
	db.SetMaxIdleConns(10) // 最大连接数据
	Conn = db

}

func runCache() error {
	//strSql := "SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE table_name='person';"
	//strSql := "SELECT * FROM person;"
	strSql := "SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE table_name='person' AND COLUMN_NAME='gender23';"
	rows, err := Conn.Query(strSql)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	exist := rows.Next()

	if exist {
		return nil
	}

	addStr := "ALTER TABLE person ADD COLUMN gender23 text;"
	_, err = Conn.Exec(addStr)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	return nil
}
