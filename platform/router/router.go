package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"goblog/platform/authenticator"
	"goblog/platform/middleware"
	"goblog/web/app/callback"
	"goblog/web/app/login"
	"goblog/web/app/logout"
	"goblog/web/app/user"
	"goblog/web/app/posts"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator, store sessions.Store) *gin.Engine {
	router := gin.Default()
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
    router.Use(sessions.Sessions("mysession", store))

	router.Static("/assets", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/logout", logout.Handler)
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/posts", middleware.IsAuthenticated, posts.ListPosts)
	router.GET("/posts/new", middleware.IsAuthenticated, posts.GetPostForm)
	router.POST("/posts/new", middleware.IsAuthenticated, posts.CreatePost)
    router.GET("/incr", func(c *gin.Context) {
        session := sessions.Default(c)
        var count int
        v := session.Get("count")
        if v == nil {
          count = 0
        } else {
          count = v.(int)
          count++
        }
        session.Set("count", count)
        session.Save()
        c.JSON(200, gin.H{"count": count})
    })

	return router
}

