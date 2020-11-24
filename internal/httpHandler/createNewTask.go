package httpHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "strconv"
	"strings"

	"AuthenticatedCRUD/model"
	"AuthenticatedCRUD/pkg/DButil"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

/*
AuthorizationヘッダのJWTをパース→nameに紐付いたパスワードのハッシュで検証→TitleとStateをJSONで受け取る→登録
*/

func CreateNewTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var (
		read io.Reader = r.Body
		task model.Task
		user model.User
	)

	db := DButil.GetClient()

	json.NewDecoder(read).Decode(&task)

	name := params.ByName("name")
	result := db.Where("name = ?", name).First(&user)
	if result.Error != nil {
		io.WriteString(w, "\"error\":\""+result.Error.Error()+"\"")
		return
	}

	if len(strings.Split(r.Header.Get("Authorization"), "Bearer ")) == 1 {
		io.WriteString(w, "\"error\":\"Token is required\"")
		return
	}
	requestToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1] //Authorizationヘッダからトークン部分だけを抜き出す．
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() == "HS256" {
			return []byte(user.Password), nil
		} else {
			return nil, fmt.Errorf("Unexpected signing method %s", token.Method.Alg())
		}
	})
	if err != nil {
		io.WriteString(w, "\"error\":\""+err.Error()+"\"")
		return
	}
	task.OwnerID, _ = token.Claims.(jwt.MapClaims)["sub"].(uint)

	if task.Title == "" {
		io.WriteString(w, "\"error\":\"Name is empty\"")
		return
	}
	if task.Status == 0 {
		task.Status = model.READY
	}

	result = db.Create(&task)
	if result.Error != nil {
		io.WriteString(w, "\"error\":\""+result.Error.Error()+"\"")
		return
	}

}
