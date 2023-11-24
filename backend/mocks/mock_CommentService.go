// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "backend/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockCommentService is an autogenerated mock type for the CommentService type
type MockCommentService struct {
	mock.Mock
}

type MockCommentService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommentService) EXPECT() *MockCommentService_Expecter {
	return &MockCommentService_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0
func (_m *MockCommentService) Add(_a0 *entity.CommentBase) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.CommentBase) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommentService_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockCommentService_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 *entity.CommentBase
func (_e *MockCommentService_Expecter) Add(_a0 interface{}) *MockCommentService_Add_Call {
	return &MockCommentService_Add_Call{Call: _e.mock.On("Add", _a0)}
}

func (_c *MockCommentService_Add_Call) Run(run func(_a0 *entity.CommentBase)) *MockCommentService_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.CommentBase))
	})
	return _c
}

func (_c *MockCommentService_Add_Call) Return(_a0 error) *MockCommentService_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommentService_Add_Call) RunAndReturn(run func(*entity.CommentBase) error) *MockCommentService_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MockCommentService) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommentService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCommentService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *MockCommentService_Expecter) Delete(id interface{}) *MockCommentService_Delete_Call {
	return &MockCommentService_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockCommentService_Delete_Call) Run(run func(id uint)) *MockCommentService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockCommentService_Delete_Call) Return(_a0 error) *MockCommentService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommentService_Delete_Call) RunAndReturn(run func(uint) error) *MockCommentService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Edit provides a mock function with given fields: _a0
func (_m *MockCommentService) Edit(_a0 *entity.CommentBase) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Edit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.CommentBase) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommentService_Edit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Edit'
type MockCommentService_Edit_Call struct {
	*mock.Call
}

// Edit is a helper method to define mock.On call
//   - _a0 *entity.CommentBase
func (_e *MockCommentService_Expecter) Edit(_a0 interface{}) *MockCommentService_Edit_Call {
	return &MockCommentService_Edit_Call{Call: _e.mock.On("Edit", _a0)}
}

func (_c *MockCommentService_Edit_Call) Run(run func(_a0 *entity.CommentBase)) *MockCommentService_Edit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.CommentBase))
	})
	return _c
}

func (_c *MockCommentService_Edit_Call) Return(_a0 error) *MockCommentService_Edit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommentService_Edit_Call) RunAndReturn(run func(*entity.CommentBase) error) *MockCommentService_Edit_Call {
	_c.Call.Return(run)
	return _c
}

// LoadAll provides a mock function with given fields: media_id
func (_m *MockCommentService) LoadAll(media_id uint) (*[]entity.CommentView, error) {
	ret := _m.Called(media_id)

	if len(ret) == 0 {
		panic("no return value specified for LoadAll")
	}

	var r0 *[]entity.CommentView
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*[]entity.CommentView, error)); ok {
		return rf(media_id)
	}
	if rf, ok := ret.Get(0).(func(uint) *[]entity.CommentView); ok {
		r0 = rf(media_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.CommentView)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(media_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommentService_LoadAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadAll'
type MockCommentService_LoadAll_Call struct {
	*mock.Call
}

// LoadAll is a helper method to define mock.On call
//   - media_id uint
func (_e *MockCommentService_Expecter) LoadAll(media_id interface{}) *MockCommentService_LoadAll_Call {
	return &MockCommentService_LoadAll_Call{Call: _e.mock.On("LoadAll", media_id)}
}

func (_c *MockCommentService_LoadAll_Call) Run(run func(media_id uint)) *MockCommentService_LoadAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockCommentService_LoadAll_Call) Return(_a0 *[]entity.CommentView, _a1 error) *MockCommentService_LoadAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommentService_LoadAll_Call) RunAndReturn(run func(uint) (*[]entity.CommentView, error)) *MockCommentService_LoadAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommentService creates a new instance of MockCommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommentService {
	mock := &MockCommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}