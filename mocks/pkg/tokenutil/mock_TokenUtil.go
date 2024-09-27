// Code generated by mockery v2.46.0. DO NOT EDIT.

package tokenutil

import (
	domain "github.com/prajnasatryass/go-clean-arch-example/internal/user/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockTokenUtil is an autogenerated mock type for the TokenUtil type
type MockTokenUtil struct {
	mock.Mock
}

type MockTokenUtil_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTokenUtil) EXPECT() *MockTokenUtil_Expecter {
	return &MockTokenUtil_Expecter{mock: &_m.Mock}
}

// CreateAccessToken provides a mock function with given fields: user, secret, ttl
func (_m *MockTokenUtil) CreateAccessToken(user *domain.User, secret string, ttl int) (string, error) {
	ret := _m.Called(user, secret, ttl)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) (string, error)); ok {
		return rf(user, secret, ttl)
	}
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, ttl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTokenUtil_CreateAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccessToken'
type MockTokenUtil_CreateAccessToken_Call struct {
	*mock.Call
}

// CreateAccessToken is a helper method to define mock.On call
//   - user *domain.User
//   - secret string
//   - ttl int
func (_e *MockTokenUtil_Expecter) CreateAccessToken(user interface{}, secret interface{}, ttl interface{}) *MockTokenUtil_CreateAccessToken_Call {
	return &MockTokenUtil_CreateAccessToken_Call{Call: _e.mock.On("CreateAccessToken", user, secret, ttl)}
}

func (_c *MockTokenUtil_CreateAccessToken_Call) Run(run func(user *domain.User, secret string, ttl int)) *MockTokenUtil_CreateAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.User), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockTokenUtil_CreateAccessToken_Call) Return(_a0 string, _a1 error) *MockTokenUtil_CreateAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTokenUtil_CreateAccessToken_Call) RunAndReturn(run func(*domain.User, string, int) (string, error)) *MockTokenUtil_CreateAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRefreshToken provides a mock function with given fields: user, secret, ttl
func (_m *MockTokenUtil) CreateRefreshToken(user *domain.User, secret string, ttl int) (string, error) {
	ret := _m.Called(user, secret, ttl)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) (string, error)); ok {
		return rf(user, secret, ttl)
	}
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, ttl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTokenUtil_CreateRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRefreshToken'
type MockTokenUtil_CreateRefreshToken_Call struct {
	*mock.Call
}

// CreateRefreshToken is a helper method to define mock.On call
//   - user *domain.User
//   - secret string
//   - ttl int
func (_e *MockTokenUtil_Expecter) CreateRefreshToken(user interface{}, secret interface{}, ttl interface{}) *MockTokenUtil_CreateRefreshToken_Call {
	return &MockTokenUtil_CreateRefreshToken_Call{Call: _e.mock.On("CreateRefreshToken", user, secret, ttl)}
}

func (_c *MockTokenUtil_CreateRefreshToken_Call) Run(run func(user *domain.User, secret string, ttl int)) *MockTokenUtil_CreateRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.User), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockTokenUtil_CreateRefreshToken_Call) Return(_a0 string, _a1 error) *MockTokenUtil_CreateRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTokenUtil_CreateRefreshToken_Call) RunAndReturn(run func(*domain.User, string, int) (string, error)) *MockTokenUtil_CreateRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTokenUtil creates a new instance of MockTokenUtil. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTokenUtil(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTokenUtil {
	mock := &MockTokenUtil{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
