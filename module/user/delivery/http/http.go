package http

import (
	"fmt"
	"strconv"
	"test/model"
	"test/module/user"
	"test/module/user/delivery"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserHttpHandler struct {
	service user.Service
}

func NewUserHttpHandler(engine *gin.Engine, service user.Service) delivery.UserHandler {
	var handler = &UserHttpHandler{
		service: service,
	}

	v1 := engine.Group("/v1/user")
	v1.GET("", handler.GetUserList)
	v1.GET("/:id", handler.GetUser)
	v1.POST("", handler.CreateUser)
	v1.PUT("/:id", handler.UpdateUser)
	v1.PATCH("/:id", handler.ModifyUser)
	v1.DELETE("/:id", handler.DeleteUser)

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

	if result, err := handler.service.GetUserList(data); err != nil {
		log.Error(err)
	} else {
		c.JSON(200, result)
	}
}

func (handler *UserHttpHandler) GetUser(c *gin.Context) {
	var data = new(model.User)

	uid, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	data.ID = uint(uid)
	if result, err := handler.service.GetUser(data); err != nil {
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
		c.JSON(500, "Internal error!")
		return
	}

	fmt.Println(data)

	if _, err := handler.service.CreateUser(data); err != nil {
		c.JSON(500, "Internal error!")
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

	if data, err = handler.service.GetUser(data); err != nil {
		log.Error(err)
	}

	if err = c.ShouldBind(&data); err != nil || data == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if _, err := handler.service.UpdateUser(data); err != nil {
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

	if _, err := handler.service.ModifyUser(data, *updateData); err != nil {
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

	if err := handler.service.DeleteUser(data); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "delete success")
	}
}
