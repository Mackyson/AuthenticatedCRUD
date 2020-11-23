package httpHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"AuthenticatedCRUD/model"
	"AuthenticatedCRUD/pkg/DButil"
	"AuthenticatedCRUD/pkg/crypto"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

/*
NameとPasswordをJSONで受け取る→Passwordハッシュ化→ハッシュを秘密鍵にしたJWT生成→JSONで返す
*/

const STRETCH_NUM = 5

func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var (
		read io.Reader = r.Body
		user model.User
	)

	db := DButil.GetClient()

	json.NewDecoder(read).Decode(&user)

	if user.Name == "" {
		io.WriteString(w, "\"error\":\"Name is empty\"")
		return
	}
	if user.Password == "" { //パスワードのバリデーションはもっと長さとか記号とか大文字小文字とかいろいろ凝るべきだが，めんどいのでなしとする．
		io.WriteString(w, "\"error\":\"Password is empty\"")
		return
	}

	user.Password = string(crypto.HashPassword([]byte(user.Password), STRETCH_NUM))

	result := db.Create(&user)
	if result.Error != nil {
		io.WriteString(w, "\"error\":\""+result.Error.Error()+"\"")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Subject:  fmt.Sprint(user.ID),
	})
	s, err := token.SignedString([]byte(user.Password))
	if err != nil {
		io.WriteString(w, "\"error\":\""+result.Error.Error()+"\"")
		return
	}

	io.WriteString(w, "\"jwt\":\""+s+"\"")

}
