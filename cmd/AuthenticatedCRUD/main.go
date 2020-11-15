package main

import (
	"log"
	"net/http"

	// "AuthenticatedCRUD/internal/handler"
	// "AuthenticatedCRUD/internal/storage"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	// storage.Migrate()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	router := httprouter.New() //API部分
	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.Write([]byte("OK!"))
	})
	// router.GET("/api/users", handler.GetUserlist) //全ユーザの提示
	// router.POST("/api/signup", handler.SignUp)    //ユーザ登録
	// router.GET("/api/signin", handler.SignIn)     //ログイン
	// //以下，認証必要
	// router.GET("/api/:name", handler.GetUserTaskList)       //全タスクの提示
	// router.GET("api/:name/:id", handler.GetUserTaskDetails) //タスクの詳細
	// router.DELETE("/api/:name/:id", handler.DeleteUserTask) //タスクの削除
	// router.PUT("/api/:name/:id", handler.UpdateUserTask)    //タスクの変更

	static := httprouter.New()
	static.ServeFiles("/*filepath", http.Dir("./webpage/static/"))
	router.NotFound = static //APIのURLとマッチせずに，not foundとなったときにファイルをサーブする

	handler := c.Handler(router) //CORSオプションの設定
	log.Println("Listen on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
