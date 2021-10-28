package http

import (
	"shopCart/config"
	"shopCart/model"
	"shopCart/middleware/auth"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initUserRoutes(api *gin.RouterGroup, conf *config.Config) {
	users := api.Group("/user")
	{
		users.POST("/signup", handler.CreateUser)
		users.POST("/login", handler.Login)
		users.POST("/logout", handler.Logout)

		authenticated := users.Use(
			cors.New(auth.CorsConfig(conf)),
			auth.AuthRequired,
		)
		{
			authenticated.GET("", handler.GetUserList)
			authenticated.GET("/:id", handler.GetUser)
			authenticated.PUT("/:id", handler.UpdateUser)
			authenticated.PATCH("/:id", handler.ModifyUser)
			authenticated.DELETE("/:id", handler.DeleteUser)
		}
	}
}

func (handler *Handler) GetUserList(c *gin.Context) {
	var data = map[string]interface{}{}

	// query url string
	if in, isExist := c.GetQuery("username"); isExist {
		data["username"] = in
	}
	if in, isExist := c.GetQuery("password"); isExist {
		data["password"] = in
	}

	if result, err := handler.services.Users.GetUserList(data); err != nil {
		log.Error(err)
	} else {
		c.JSON(200, result)
	}
}

func (handler *Handler) GetUser(c *gin.Context) {
	var data = new(model.User)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)
	if result, err := handler.services.Users.GetUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, result)
	}
}

func (handler *Handler) CreateUser(c *gin.Context) {
	var data = new(model.User)

	if err := c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if _, err := handler.services.Users.CreateUser(data); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	} else {
		c.JSON(200, "create user success")
	}

}

func (handler *Handler) UpdateUser(c *gin.Context) {
	var (
		data = new(model.User)
		err  error
	)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)

	if data, err = handler.services.Users.GetUser(data); err != nil {
		log.Error(err)
	}

	if err = c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if _, err := handler.services.Users.UpdateUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "update user success")
	}
}

func (handler *Handler) ModifyUser(c *gin.Context) {
	var (
		data       = new(model.User)
		updateData = new(map[string]interface{})
	)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)

	if err := c.ShouldBind(&updateData); err != nil || updateData == nil {
		log.Error(err)
		c.JSON(500, "internal error!")
		return
	}

	if _, err := handler.services.Users.ModifyUser(data, *updateData); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "modify user success")
	}
}

func (handler *Handler) DeleteUser(c *gin.Context) {
	var data = new(model.User)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)

	if err := handler.services.Users.DeleteUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "delete success")
	}
}

func (handler *Handler) Login(c *gin.Context) {
	var data = new(model.User)
	if err := c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if _, err := handler.services.Users.GetUser(data); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if err := auth.SaveToRedis(c); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{
		"isLogined": true,
	})
}

func (handler *Handler) Logout(c *gin.Context) {
	if err := auth.DeleteFromRedis(c); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{
		"isLogined": false,
	})
}
