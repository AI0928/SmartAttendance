package main

import (
	"app/database"
	"app/handler"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    // Echoインスタンスを作成
    e := echo.New()

    // httpリクエストの情報をログに表示
	e.Use(middleware.Logger())
	// パニックを回復し、スタックトレースを表示
	e.Use(middleware.Recover())

	// CORS対策
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "127.0.0.1:3000", "http://localhost:8080"}, // ドメイン
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	database.DBConnect()
    // ルートを設定（第一引数にエンドポイント、第二引数にハンドラーを指定）
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.GET("/users", handler.GetUsers)
	api := e.Group("/api")
    api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: handler.Config.SigningKey,})) // /api 下はJWTの認証が必要
	api.GET("/user", handler.GetMe)


    // サーバーをポート番号8080で起動
    e.Logger.Fatal(e.Start(":8080"))
}


