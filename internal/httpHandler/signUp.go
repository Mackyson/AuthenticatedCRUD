package httpHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"AuthenticatedCRUD/model"
	"AuthenticatedCRUD/pkg/DButil"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var (
		read io.Reader = r.Body
		user model.User
	)

	json.NewDecoder(read).Decode(&user)

	fmt.Printf("%+v\n%+v", user, &model.User{}) //dbg
	db := DButil.GetClient()

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
