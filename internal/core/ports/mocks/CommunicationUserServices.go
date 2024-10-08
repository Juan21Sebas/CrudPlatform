// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	repository "CrudPlatform/internal/core/domain/repository"

	users "CrudPlatform/internal/core/domain/repository/model/users"
)

// CommunicationUserServices is an autogenerated mock type for the CommunicationUserServices type
type CommunicationUserServices struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, request
func (_m *CommunicationUserServices) CreateUser(ctx *gin.Context, request *users.User) (*repository.Response, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *repository.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.User) (*repository.Response, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.User) *repository.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, *users.User) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, request
func (_m *CommunicationUserServices) DeleteUser(ctx *gin.Context, request *users.DeleteUser) (*repository.Response, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 *repository.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.DeleteUser) (*repository.Response, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.DeleteUser) *repository.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, *users.DeleteUser) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUser provides a mock function with given fields: ctx, request
func (_m *CommunicationUserServices) SelectUser(ctx *gin.Context, request *users.GetUser) (*repository.Response, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SelectUser")
	}

	var r0 *repository.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.GetUser) (*repository.Response, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.GetUser) *repository.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, *users.GetUser) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, request
func (_m *CommunicationUserServices) UpdateUser(ctx *gin.Context, request *users.UpdateUser) (*repository.Response, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *repository.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.UpdateUser) (*repository.Response, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.UpdateUser) *repository.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, *users.UpdateUser) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCommunicationUserServices creates a new instance of CommunicationUserServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommunicationUserServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommunicationUserServices {
	mock := &CommunicationUserServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
