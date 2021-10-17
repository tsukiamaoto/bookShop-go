package http

import (
	Config "shopCart/config"
	"shopCart/middleware/auth"
	"shopCart/model"
	"shopCart/module/user"
	"shopCart/module/user/delivery"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

type UserHttpHandler struct {
	Service user.Service
}

var store *sessions.CookieStore

func init() {
	config := Config.LoadConfig()
	// create new session id and save into "store" global viariable
	store = sessions.NewCookieStore([]byte(config.SessionKey))
	store.MaxAge(86400 * 7)
}

func NewUserHttpHandler(engine *gin.Engine, service user.Service) delivery.UserHandler {
	var handler = &UserHttpHandler{
		Service: service,
	}

	v1 := engine.Group("/v1")
	{
		userApi := v1.Group("/user")
		userApi.Use(auth.AuthRequired)
		userApi.GET("", handler.GetUserList)
		userApi.GET("/:id", handler.GetUser)
		userApi.POST("", handler.CreateUser)
		userApi.PUT("/:id", handler.UpdateUser)
		userApi.PATCH("/:id", handler.ModifyUser)
		userApi.DELETE("/:id", handler.DeleteUser)
	}
	{
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

	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	session.Values["auth"] = true
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "logged in successfully!")
}

func (handler *UserHttpHandler) Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	session.Values["auth"] = nil
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "logged out successfully!")
}
