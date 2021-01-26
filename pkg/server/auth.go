package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	hydra "github.com/ory/hydra-client-go/client"
	log "github.com/sirupsen/logrus"

	"wedding/ent"
	"wedding/ent/backroomuser"
)

type Auth struct {
	database *ent.Client
	router   *gin.RouterGroup
	admin    *hydra.OryHydra
	public   *hydra.OryHydra
	store    redis.Store
}

func RegisterAuth(database *ent.Client, g *gin.RouterGroup, store redis.Store) *Auth {
	api := &Auth{
		database: database,
		router:   g,
		store:    store,
	}

	g.POST("/login", api.login)
	g.GET("/logout", api.logout)

	return api
}

func (*Auth) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		timeout := session.Get("timeout")

		if username == nil || timeout == nil || (time.Unix(timeout.(int64), 0).Before(time.Now())) {
			c.AbortWithStatus(http.StatusForbidden)
		}

		session.Set("timeout", time.Now().Add(time.Minute*30))
	}
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (api *Auth) login(c *gin.Context) {
	session := sessions.Default(c)

	credentials := &login{}

	err := c.MustBindWith(credentials, binding.JSON)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := api.database.BackroomUser.Query().
		Where(backroomuser.Username(credentials.Username)).
		Only(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	session.Set("username", user.Username)
	session.Set("timeout", time.Now().Add(time.Minute*30).Unix())
	err = session.Save()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"user":    user,
		"timeout": (time.Minute * 30).Milliseconds(),
	})
}

func (api *Auth) logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("username")
	session.Delete("timeout")
	err := session.Save()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	c.Status(200)
}
