// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "test/model"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: in
func (_m *Service) CreateUser(in *model.User) (*model.User, error) {
	ret := _m.Called(in)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: in
func (_m *Service) DeleteUser(in *model.User) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: in
func (_m *Service) GetUser(in *model.User) (*model.User, error) {
	ret := _m.Called(in)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserList provides a mock function with given fields: _a0
func (_m *Service) GetUserList(_a0 map[string]interface{}) ([]*model.User, error) {
	ret := _m.Called(_a0)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(map[string]interface{}) []*model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyUser provides a mock function with given fields: in, data
func (_m *Service) ModifyUser(in *model.User, data map[string]interface{}) (*model.User, error) {
	ret := _m.Called(in, data)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User, map[string]interface{}) *model.User); ok {
		r0 = rf(in, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, map[string]interface{}) error); ok {
		r1 = rf(in, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: in
func (_m *Service) UpdateUser(in *model.User) (*model.User, error) {
	ret := _m.Called(in)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}