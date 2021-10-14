package auth

import (
	"net/http"
	Config "test/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

var store *sessions.CookieStore

func init() {
	config := Config.LoadConfig()
	// create new session id and save into "store" global viariable
	store = sessions.NewCookieStore([]byte(config.SessionKey))
	store.MaxAge(86400 * 7)
}

func AuthRequired(c *gin.Context) {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}
	// if auth is nil return error message
	if session.Values["auth"] == nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}