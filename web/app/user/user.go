package user

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"goblog/platform/database"
)

type User struct {
	Id       int;
	OauthSub string `db:"oauth_sub"`;
    Picture  string;
}

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	user := User{}
	sub := profile.(map[string]interface{})["sub"].(string)
	picture := profile.(map[string]interface{})["picture"].(string)
	err := database.Db.Get(&user, "SELECT id FROM users WHERE oauth_sub=$1", sub)
    user.OauthSub = sub
    user.Picture = picture
	if err != nil {
		log.Println(err)
		ctx.HTML(http.StatusInternalServerError, "InternalError.html", nil)
		return
	}
	log.Println(user)

	ctx.HTML(http.StatusOK, "user.html", user)
}

