// Code generated by mockery v2.46.0. DO NOT EDIT.

package hasher

import mock "github.com/stretchr/testify/mock"

// MockHasher is an autogenerated mock type for the Hasher type
type MockHasher struct {
	mock.Mock
}

type MockHasher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHasher) EXPECT() *MockHasher_Expecter {
	return &MockHasher_Expecter{mock: &_m.Mock}
}

// HashPassword provides a mock function with given fields: password
func (_m *MockHasher) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHasher_HashPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashPassword'
type MockHasher_HashPassword_Call struct {
	*mock.Call
}

// HashPassword is a helper method to define mock.On call
//   - password string
func (_e *MockHasher_Expecter) HashPassword(password interface{}) *MockHasher_HashPassword_Call {
	return &MockHasher_HashPassword_Call{Call: _e.mock.On("HashPassword", password)}
}

func (_c *MockHasher_HashPassword_Call) Run(run func(password string)) *MockHasher_HashPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHasher_HashPassword_Call) Return(_a0 string, _a1 error) *MockHasher_HashPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHasher_HashPassword_Call) RunAndReturn(run func(string) (string, error)) *MockHasher_HashPassword_Call {
	_c.Call.Return(run)
	return _c
}

// MatchPassword provides a mock function with given fields: password, hash
func (_m *MockHasher) MatchPassword(password string, hash string) bool {
	ret := _m.Called(password, hash)

	if len(ret) == 0 {
		panic("no return value specified for MatchPassword")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockHasher_MatchPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MatchPassword'
type MockHasher_MatchPassword_Call struct {
	*mock.Call
}

// MatchPassword is a helper method to define mock.On call
//   - password string
//   - hash string
func (_e *MockHasher_Expecter) MatchPassword(password interface{}, hash interface{}) *MockHasher_MatchPassword_Call {
	return &MockHasher_MatchPassword_Call{Call: _e.mock.On("MatchPassword", password, hash)}
}

func (_c *MockHasher_MatchPassword_Call) Run(run func(password string, hash string)) *MockHasher_MatchPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockHasher_MatchPassword_Call) Return(_a0 bool) *MockHasher_MatchPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHasher_MatchPassword_Call) RunAndReturn(run func(string, string) bool) *MockHasher_MatchPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHasher creates a new instance of MockHasher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHasher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHasher {
	mock := &MockHasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
