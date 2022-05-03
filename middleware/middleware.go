package middleware

import (
	"net/http"
	"tsukiamaoto/bookShop-go/config"
	"tsukiamaoto/bookShop-go/redis"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	log "github.com/sirupsen/logrus"
)

var store *redisstore.RedisStore

func init() {
	var (
		err error
	)

	// set redis configuration
	store, err = redisstore.NewRedisStore(redis.GetRDB().Context(), redis.GetRDB())
	if err != nil {
		log.Error("failed to create redis store: ", err)
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		Path: "/",
		// Domain: "localhost",
		MaxAge: 86400 * 7, // 7 days
	})

}

func CorsConfig(conf *config.Config) cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = conf.AllowOrigins
	// corsConf.AllowCredentials = true
	// corsConf.AllowAllOrigins = true
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}

	return corsConf
}

func AuthRequired(c *gin.Context) {
	session, err := store.Get(c.Request, "session-key")
	if err != nil {
		log.Error("Failded to get session, reason is:", err)
		c.JSON(500, err.Error())
		return
	}
	// if auth is nil return error message
	if session.Values["auth"] == nil {
		log.Error("Failded to get auth from session, reason is:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}

func GetAuth(c *gin.Context) (bool, error) {
	var isLogined bool

	session, err := store.Get(c.Request, "session-key")
	if err != nil {
		log.Error("Failed to get session, reason is :", err)
		return false, err
	}

	if session.Values["auth"] == true {
		isLogined = true
	}

	return isLogined, nil
}

func SetAuth(c *gin.Context) error {
	session, err := store.Get(c.Request, "session-key")
	if err != nil {
		log.Error("Failed to get session, reason is :", err)
		return err
	}

	session.Values["auth"] = true
	if err = session.Save(c.Request, c.Writer); err != nil {
		log.Error("Failed to save session, reason is :", err)
		return err
	}

	return nil
}

func DeleteAuth(c *gin.Context) error {
	session, err := store.Get(c.Request, "session-key")
	if err != nil {
		log.Error("Failed to get session, reason is :", err)
		return err
	}

	// redis delete data when Maxage <= 0
	session.Options.MaxAge = -1
	if err = session.Save(c.Request, c.Writer); err != nil {
		log.Error("Failed to delete session, reason is :", err)
		return err
	}

	return nil
}

func GetUserId(c *gin.Context) (uint, error) {
	session, err := store.Get(c.Request, "session-key")
	if err != nil {
		log.Error("Failed to get session, reason is :", err)
		return 0, err
	}

	val := session.Values["userId"]

	userIdOfString, ok := val.(string)
	if !ok {
		log.Error("Failed to get session value by key: 'userId'")
	}

	uid64, _ := strconv.ParseUint(userIdOfString, 10, 64)
	userId := uint(uid64)

	return userId, err
}

func SaveToRedis(c *gin.Context, key, value string) error {
	session, err := store.Get(c.Request, "session-key")
	if err != nil {
		log.Error("Failed to get session, reason is :", err)
		return err
	}

	session.Values[key] = value
	if err = session.Save(c.Request, c.Writer); err != nil {
		log.Error("Failed to save session, reason is :", err)
		return err
	}

	return nil
}
