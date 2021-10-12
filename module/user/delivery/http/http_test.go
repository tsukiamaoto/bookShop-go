package http_test

import (
	"strconv"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	

	mocks "test/mocks/module/user"
	"test/model"
	userHttpHandlerDelivery "test/module/user/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserList(t *testing.T) {
	mockUsers := []*model.User{
		&model.User{
			ID:       uint(1),
			Username: "123",
			Password: "123",
		},
		&model.User{
			ID:       uint(2),
			Username: "123",
			Password: "123",
		},
	}
	mockUsersMarshal, _ := json.Marshal(mockUsers)
	mockService := new(mocks.Service)

	mockService.
		On("GetUserList", mock.Anything).
		Return(mockUsers, nil)

	handler := userHttpHandlerDelivery.UserHttpHandler{
		Service: mockService,
	}

	engine := gin.Default()
	v1 := engine.Group("/v1/user")
	v1.GET("", handler.GetUserList)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(mockUsersMarshal), w.Body.String())
}

func TestGetUser(t *testing.T) {
	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}
	mockUserMarshal, _ := json.Marshal(mockUser)
	mockService := new(mocks.Service)

	mockService.
		On("GetUser", mock.AnythingOfType("*model.User")).
		Return(mockUser, nil)

	handler := userHttpHandlerDelivery.UserHttpHandler{
		Service: mockService,
	}

	engine := gin.Default()
	v1 := engine.Group("/v1/user")
	v1.GET("/:id", handler.GetUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(mockUserMarshal), w.Body.String())
}

func TestCreateUser(t *testing.T) {
	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}
	mockService := new(mocks.Service)

	mockService.
		On("CreateUser", mock.AnythingOfType("*model.User")).
		Return(mockUser, nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*model.User)
			arg.ID = uint(1)
		})

	handler := userHttpHandlerDelivery.UserHttpHandler{
		Service: mockService,
	}

	engine := gin.Default()
	v1 := engine.Group("/v1/user")
	v1.POST("", handler.CreateUser)

	data := url.Values{}
	data.Set("username", "123")
	data.Set("password", "123")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/user", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// func TestUpdateUser(t *testing.T) {
// 	mockUser := &model.User{
// 		ID:       uint(1),
// 		Username: "456",
// 		Password: "456",
// 	}
// 	mockService := new(mocks.Service)

// 	mockService.
// 		On("UpdateUser", mock.AnythingOfType("*model.User")).
// 		Return(mockUser, nil)

// 	handler := userHttpHandlerDelivery.UserHttpHandler{
// 		Service: mockService,
// 	}

// 	engine := gin.Default()
// 	v1 := engine.Group("/v1/user")
// 	v1.PUT("/:id", handler.UpdateUser)

// 	data := url.Values{}
// 	data.Set("username", "123")
// 	data.Set("password", "123")

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPut, "/v1/user/1", strings.NewReader(data.Encode()))
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
// 	engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }
