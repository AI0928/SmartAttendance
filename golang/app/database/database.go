package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 外部参照可能なDB変数を定義
var DB *gorm.DB

// DBを起動させる
func DBConnect() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	container_name := "smart_attendance_db"
	database_name := os.Getenv("MYSQL_DATABASE")
	// [ユーザ名]:[パスワード]@tcp([DBコンテナ名])/[データベース名]?charset=[文字コード]
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True`, 
		user, password, container_name, database_name)
	// DBへの接続を行う
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	// エラーが発生した場合、エラー内容を表示
	if err != nil {
		fmt.Println(err)
	}
	// 接続に成功した場合、「db connected!!」と表示する
	fmt.Println("db connected!!")
}