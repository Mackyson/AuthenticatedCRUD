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
NameとPasswordをJSONで受け取る→Password比較→ハッシュを秘密鍵にしたJWT生成→JSONで返す
*/
func SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var (
		read      io.Reader  = r.Body
		userInput model.User //FIXME フロントエンドからの入力は別の型を作ったほうがいい？どうなんでしょう
		userDB    model.User
	)

	db := DButil.GetClient()

	json.NewDecoder(read).Decode(&userInput)

	result := db.Where("name = ?", userInput.Name).First(&userDB)
	if result.Error != nil {
		io.WriteString(w, "\"error\":\""+result.Error.Error()+"\"")
		return
	}
	if !crypto.IsValidPassword([]byte(userDB.Password), []byte(userInput.Password)) {
		http.Error(w, fmt.Sprintf("%d Unauthorized", http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Subject:  fmt.Sprint(userDB.ID),
	})
	s, err := token.SignedString([]byte(userDB.Password))
	if err != nil {
		io.WriteString(w, "\"error\":\""+result.Error.Error()+"\"")
		return
	}

	io.WriteString(w, "\"jwt\":\""+s+"\"")

}
