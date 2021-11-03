package http

import (
	"shopCart/config"
	"shopCart/middleware"
	"shopCart/model"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initUserRoutes(api *gin.RouterGroup, conf *config.Config) {
	users := api.Group("/user")
	{
		users.POST("/signup", handler.Signup)
		users.POST("/login", handler.Login)
		users.POST("/logout", handler.Logout)

		authenticated := users.Use(
			cors.New(middleware.CorsConfig(conf)),
			middleware.AuthRequired,
		)
		{
			authenticated.GET("", handler.GetUserList)
			authenticated.GET("/:id", handler.GetUser)
			authenticated.PUT("/:id", handler.UpdateUser)
			authenticated.DELETE("/:id", handler.DeleteUser)
		}
	}
}

func (handler *Handler) GetUserList(c *gin.Context) {
	if users, err := handler.services.Users.GetUserList(); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, users)
	}
}

func (handler *Handler) GetUser(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	userId := uint(uid64)

	if user, err := handler.services.Users.GetUserById(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, user)
	}
}

func (handler *Handler) UpdateUser(c *gin.Context) {
	var (
		user = new(model.User)
		err  error
	)

	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	user.ID = uint(uid64)

	if err = c.ShouldBind(&user); err != nil || user == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if _, err = handler.services.Users.UpdateUser(user); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "Updated user successfully!")
	}
}

func (handler *Handler) DeleteUser(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	userId := uint(uid64)

	if err := handler.services.Users.DeleteUser(userId); err != nil {
		log.Error(err)
		c.JSON(500, "internal error!")
	} else {
		c.JSON(200, "Deleted user successfully")
	}
}

func (handler *Handler) Login(c *gin.Context) {
	var user = new(model.User)
	if err := c.ShouldBind(&user); err != nil || user == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if _, err := handler.services.Users.GetUser(user); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if err := middleware.SetAuth(c); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{
		"isLogined": true,
	})
}

func (handler *Handler) Logout(c *gin.Context) {
	if err := middleware.DeleteAuth(c); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{
		"isLogined": false,
	})
}

// create user, and then create order and cart instance association with user
func (handler *Handler) Signup(c *gin.Context) {
	var user = new(model.User)
	
	if err := c.ShouldBind(&user); err != nil || user == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	//create user info
	var (
		newUser *model.User
		err     error
	)
	if newUser, err = handler.services.Users.CreateUser(user); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}
	userId := newUser.ID

	//create cart instance and associate with user
	if err := handler.services.Carts.CreateCartWithUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	// creat order instance and associate with user
	if err := handler.services.Orders.CreateOrderWithUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	// save userID into session
	uid64 := uint64(userId)
	strUserId := strconv.FormatUint(uid64, 10)
	if err := middleware.SaveToRedis(c, "userId", strUserId); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "create user successfully!")
}
