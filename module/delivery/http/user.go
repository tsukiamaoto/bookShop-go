package http

import (
	"github.com/tsukiamaoto/bookShop-go/config"
	"github.com/tsukiamaoto/bookShop-go/middleware"
	"github.com/tsukiamaoto/bookShop-go/model"

	"strconv"

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

// @Summary Get User List
// @Tags User
// @Description get user list
// @ModuleID getUserList
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]model.User} "get users"
// @Failure 500 string Internal error!
// @Router /user [get]
func (handler *Handler) GetUserList(c *gin.Context) {
	var (
		users []*model.User
		err error
	)
	if users, err = handler.services.Users.GetUserList(); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	}

	c.JSON(200, dataResponse{Data: users})
}

// @Summary Get User
// @Tags User
// @Description get user by user id
// @ModuleID getUser
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dataResponse{data=model.User} "get the user"
// @Failure 500 string Internal error!
// @Router /user/:id [get]
func (handler *Handler) GetUser(c *gin.Context) {
	var (
		user *model.User
		err  error
	)
	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	userId := uint(uid64)

	if user, err = handler.services.Users.GetUserById(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, dataResponse{Data: user})
}

// @Summary Update User
// @Tags User
// @Description update user by user id
// @ModuleID updateUser
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} dataResponse{data=string} "Updated user successfully!"
// @Failure 500 string parameter error!
// @Failure 500 string Internal error
// @Router /user/:id [put]
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
		c.JSON(500, "Internal!")
		return
	}

	c.JSON(200, dataResponse{Data: "Updated user successfully!"})
}

// @Summary Delete User
// @Tags User
// @Description delete user by user id
// @ModuleID deleteUser
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dataResponse{data=string} "Deleted user successfully"
// @Failure 500 string Internal error!
// @Router /user/:id [delete]
func (handler *Handler) DeleteUser(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	userId := uint(uid64)

	if err := handler.services.Users.DeleteUser(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal!")
		return
	}

	c.JSON(200, dataResponse{Data: "Deleted user successfully"})
}

// @Summary login
// @Tags User
// @Description login
// @ModuleID login
// @Accept json
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} object{isLogined=boolean} "isLogined is true"
// @Failure 500 string error message
// @Router /user/login [post]
func (handler *Handler) Login(c *gin.Context) {
	var (
		user      = new(model.User)
		isLogined = false
		err       error
	)

	// checked whatever session-key existed or not
	if isLogined, err = middleware.GetAuth(c); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	// if session-key existed
	if isLogined {
		c.JSON(200, gin.H{
			"isLogined": isLogined,
		})
		return
	}

	// examined user format is correct
	if err = c.ShouldBind(&user); err != nil || user == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	// checked whether user existed or not
	if _, err = handler.services.Users.GetUser(user); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}
	isLogined = true

	// set auth to session
	if err = middleware.SetAuth(c); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"isLogined": isLogined,
	})
}

// @Summary logout
// @Tags User
// @Description logout by session userid
// @ModuleID logout
// @Accept json
// @Produce json
// @Success 200 {object} object{isLogined=boolean} "isLogined is false"
// @Failure 500 string error message
// @Router /user/logout [post]
func (handler *Handler) Logout(c *gin.Context) {
	if err := middleware.DeleteAuth(c); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{
		"isLogined": false,
	})
}

// @Summary signup
// @Tags User
// @Description signup
// @ModuleID signup
// @Accept json
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} dataResponse{Data:"create user successfully!"} "desc"
// @Header 200 {string} session-key "user id"
// @Failure 500 string error message
// @Router /user/signup [post]

// create user, and then create order and cart instance association with user
func (handler *Handler) Signup(c *gin.Context) {
	var (
		user = new(model.User)
		err  error
	)

	if err := c.ShouldBind(&user); err != nil || user == nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}
	// check whether user existed or not by username
	var oldUser *model.User
	if oldUser, err = handler.services.Users.GetUser(&model.User{Username: user.Username}); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	if oldUser.ID != 0 {
		c.JSON(200, dataResponse{Data: "User has existed!"})
		return
	}

	//create a new user
	var newUser *model.User
	if newUser, err = handler.services.Users.CreateUser(user); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}
	userId := newUser.ID

	// create seller instance and associate with user
	if err := handler.services.Sellers.CreateSellerWithUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

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
	strUserId := strconv.FormatUint(uint64(userId), 10)
	if err := middleware.SaveToRedis(c, "userId", strUserId); err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, dataResponse{Data: "Created user successfully!"})
}
