package posts

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"goblog/platform/database"
)

type PostList struct {
	Posts []Post
}

type Post struct {
	Id      int
	Title   string
	Created time.Time
	Author  string
}

const listPostsSql = `
SELECT
	p.id,
	p.title,
	p.created,
	u.name as author
from posts p
join users u
on u.id = p.author_id;`

func ListPosts(ctx *gin.Context) {
	posts := []Post{}
	err := database.Db.Select(&posts, listPostsSql)
	if err != nil {
		log.Printf("Error encountered %+v", err)
		ctx.HTML(http.StatusInternalServerError, "InternalError.html", nil)
		return
	}

	ctx.HTML(http.StatusOK, "posts.html", PostList{Posts: posts})
}

func GetPostForm(c *gin.Context) {
	c.HTML(http.StatusOK, "NewPostForm.html", nil)
}

type NewPost struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}

const createPostSql = `
INSERT INTO posts (title, body, author_id)
values ($1, $2, (select id from users where oauth_sub = $3));`

func CreatePost(ctx *gin.Context) {
	var new_post NewPost

	bind_err := ctx.ShouldBind(&new_post)
	if bind_err != nil {
		ctx.HTML(http.StatusBadRequest, "BadRequest.html", nil)
		return
	}

	session := sessions.Default(ctx)
	profile := session.Get("profile")

	sub := profile.(map[string]interface{})["sub"].(string)
	log.Printf("input %+v sub %+v", new_post, sub)

	_, write_err := database.Db.Exec(createPostSql, new_post.Title, new_post.Body, sub)
	if write_err != nil {
		log.Printf("Error encountered %+v", write_err)
		ctx.HTML(http.StatusInternalServerError, "InternalError.html", nil)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/posts")
}
