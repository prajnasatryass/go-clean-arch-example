// Code generated by mockery v2.46.0. DO NOT EDIT.

package domain

import (
	domain "github.com/prajnasatryass/tic-be/internal/user/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: user
func (_m *MockUserRepository) Create(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - user *domain.User
func (_e *MockUserRepository_Expecter) Create(user interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", user)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(user *domain.User)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.User))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(*domain.User) error) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByID provides a mock function with given fields: id
func (_m *MockUserRepository) DeleteByID(id uuid.UUID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_DeleteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByID'
type MockUserRepository_DeleteByID_Call struct {
	*mock.Call
}

// DeleteByID is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *MockUserRepository_Expecter) DeleteByID(id interface{}) *MockUserRepository_DeleteByID_Call {
	return &MockUserRepository_DeleteByID_Call{Call: _e.mock.On("DeleteByID", id)}
}

func (_c *MockUserRepository_DeleteByID_Call) Run(run func(id uuid.UUID)) *MockUserRepository_DeleteByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockUserRepository_DeleteByID_Call) Return(_a0 error) *MockUserRepository_DeleteByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_DeleteByID_Call) RunAndReturn(run func(uuid.UUID) error) *MockUserRepository_DeleteByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetByEmail(email string) (domain.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type MockUserRepository_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockUserRepository_Expecter) GetByEmail(email interface{}) *MockUserRepository_GetByEmail_Call {
	return &MockUserRepository_GetByEmail_Call{Call: _e.mock.On("GetByEmail", email)}
}

func (_c *MockUserRepository_GetByEmail_Call) Run(run func(email string)) *MockUserRepository_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetByEmail_Call) Return(_a0 domain.User, _a1 error) *MockUserRepository_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetByEmail_Call) RunAndReturn(run func(string) (domain.User, error)) *MockUserRepository_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: id
func (_m *MockUserRepository) GetByID(id uuid.UUID) (domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockUserRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *MockUserRepository_Expecter) GetByID(id interface{}) *MockUserRepository_GetByID_Call {
	return &MockUserRepository_GetByID_Call{Call: _e.mock.On("GetByID", id)}
}

func (_c *MockUserRepository_GetByID_Call) Run(run func(id uuid.UUID)) *MockUserRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockUserRepository_GetByID_Call) Return(_a0 domain.User, _a1 error) *MockUserRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetByID_Call) RunAndReturn(run func(uuid.UUID) (domain.User, error)) *MockUserRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// PermaDeleteByID provides a mock function with given fields: id
func (_m *MockUserRepository) PermaDeleteByID(id uuid.UUID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for PermaDeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_PermaDeleteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PermaDeleteByID'
type MockUserRepository_PermaDeleteByID_Call struct {
	*mock.Call
}

// PermaDeleteByID is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *MockUserRepository_Expecter) PermaDeleteByID(id interface{}) *MockUserRepository_PermaDeleteByID_Call {
	return &MockUserRepository_PermaDeleteByID_Call{Call: _e.mock.On("PermaDeleteByID", id)}
}

func (_c *MockUserRepository_PermaDeleteByID_Call) Run(run func(id uuid.UUID)) *MockUserRepository_PermaDeleteByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockUserRepository_PermaDeleteByID_Call) Return(_a0 error) *MockUserRepository_PermaDeleteByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_PermaDeleteByID_Call) RunAndReturn(run func(uuid.UUID) error) *MockUserRepository_PermaDeleteByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
