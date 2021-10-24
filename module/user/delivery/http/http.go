package http

import (
	"shopCart/middleware/auth"
	"shopCart/model"
	"shopCart/module/user"
	"shopCart/module/user/delivery"
	"shopCart/config"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserHttpHandler struct {
	Service user.Service
}

func NewUserHttpHandler(engine *gin.Engine, service user.Service) delivery.UserHandler {
	var handler = &UserHttpHandler{
		Service: service,
	}
	var config = config.LoadConfig()

	v1 := engine.Group("/v1")
	{
		userApi := v1.Group("/user")
		userApi.Use(cors.New(auth.CorsConfig(config)))
		userApi.Use(auth.AuthRequired)
		userApi.GET("", handler.GetUserList)
		userApi.GET("/:id", handler.GetUser)
		userApi.PUT("/:id", handler.UpdateUser)
		userApi.PATCH("/:id", handler.ModifyUser)
		userApi.DELETE("/:id", handler.DeleteUser)
	}
	{
		v1.POST("/user", handler.CreateUser)
		v1.POST("/login", handler.Login)
		v1.POST("/logout", handler.Logout)
	}

	return handler
}

func (handler *UserHttpHandler) GetUserList(c *gin.Context) {
	var data = map[string]interface{}{}

	// query url string
	if in, isExist := c.GetQuery("username"); isExist {
		data["username"] = in
	}
	if in, isExist := c.GetQuery("password"); isExist {
		data["password"] = in
	}

	if result, err := handler.Service.GetUserList(data); err != nil {
		log.Error(err)
	} else {
		c.JSON(200, result)
	}
}

func (handler *UserHttpHandler) GetUser(c *gin.Context) {
	var data = new(model.User)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)
	if result, err := handler.Service.GetUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, result)
	}
}

func (handler *UserHttpHandler) CreateUser(c *gin.Context) {
	var data = new(model.User)

	if err := c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if _, err := handler.Service.CreateUser(data); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	} else {
		c.JSON(200, "create user success")
	}

}

func (handler *UserHttpHandler) UpdateUser(c *gin.Context) {
	var (
		data = new(model.User)
		err  error
	)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)

	if data, err = handler.Service.GetUser(data); err != nil {
		log.Error(err)
	}

	if err = c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if _, err := handler.Service.UpdateUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "update user success")
	}
}

func (handler *UserHttpHandler) ModifyUser(c *gin.Context) {
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

	if _, err := handler.Service.ModifyUser(data, *updateData); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "modify user success")
	}
}

func (handler *UserHttpHandler) DeleteUser(c *gin.Context) {
	var data = new(model.User)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)

	if err := handler.Service.DeleteUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "delete success")
	}
}

func (handler *UserHttpHandler) Login(c *gin.Context) {
	var data = new(model.User)
	if err := c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if _, err := handler.Service.GetUser(data); err != nil {
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

func (handler *UserHttpHandler) Logout(c *gin.Context) {
	if err := auth.DeleteFromRedis(c); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{
		"isLogined": false,
	})
}
