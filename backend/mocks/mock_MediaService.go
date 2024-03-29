// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "backend/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockMediaService is an autogenerated mock type for the MediaService type
type MockMediaService struct {
	mock.Mock
}

type MockMediaService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMediaService) EXPECT() *MockMediaService_Expecter {
	return &MockMediaService_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0
func (_m *MockMediaService) Add(_a0 *entity.Media) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Media) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaService_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockMediaService_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 *entity.Media
func (_e *MockMediaService_Expecter) Add(_a0 interface{}) *MockMediaService_Add_Call {
	return &MockMediaService_Add_Call{Call: _e.mock.On("Add", _a0)}
}

func (_c *MockMediaService_Add_Call) Run(run func(_a0 *entity.Media)) *MockMediaService_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Media))
	})
	return _c
}

func (_c *MockMediaService_Add_Call) Return(_a0 error) *MockMediaService_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaService_Add_Call) RunAndReturn(run func(*entity.Media) error) *MockMediaService_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MockMediaService) Delete(id uint) error {
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

// MockMediaService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockMediaService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *MockMediaService_Expecter) Delete(id interface{}) *MockMediaService_Delete_Call {
	return &MockMediaService_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockMediaService_Delete_Call) Run(run func(id uint)) *MockMediaService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockMediaService_Delete_Call) Return(_a0 error) *MockMediaService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaService_Delete_Call) RunAndReturn(run func(uint) error) *MockMediaService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Edit provides a mock function with given fields: _a0
func (_m *MockMediaService) Edit(_a0 *entity.Media) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Edit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Media) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaService_Edit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Edit'
type MockMediaService_Edit_Call struct {
	*mock.Call
}

// Edit is a helper method to define mock.On call
//   - _a0 *entity.Media
func (_e *MockMediaService_Expecter) Edit(_a0 interface{}) *MockMediaService_Edit_Call {
	return &MockMediaService_Edit_Call{Call: _e.mock.On("Edit", _a0)}
}

func (_c *MockMediaService_Edit_Call) Run(run func(_a0 *entity.Media)) *MockMediaService_Edit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Media))
	})
	return _c
}

func (_c *MockMediaService_Edit_Call) Return(_a0 error) *MockMediaService_Edit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaService_Edit_Call) RunAndReturn(run func(*entity.Media) error) *MockMediaService_Edit_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields: id
func (_m *MockMediaService) Load(id uint) (*entity.MediaView, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 *entity.MediaView
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*entity.MediaView, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *entity.MediaView); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.MediaView)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMediaService_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MockMediaService_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - id uint
func (_e *MockMediaService_Expecter) Load(id interface{}) *MockMediaService_Load_Call {
	return &MockMediaService_Load_Call{Call: _e.mock.On("Load", id)}
}

func (_c *MockMediaService_Load_Call) Run(run func(id uint)) *MockMediaService_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockMediaService_Load_Call) Return(_a0 *entity.MediaView, _a1 error) *MockMediaService_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMediaService_Load_Call) RunAndReturn(run func(uint) (*entity.MediaView, error)) *MockMediaService_Load_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: _a0
func (_m *MockMediaService) Search(_a0 *entity.Filter) (*[]entity.MediaResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *[]entity.MediaResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Filter) (*[]entity.MediaResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*entity.Filter) *[]entity.MediaResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.MediaResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Filter) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMediaService_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockMediaService_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - _a0 *entity.Filter
func (_e *MockMediaService_Expecter) Search(_a0 interface{}) *MockMediaService_Search_Call {
	return &MockMediaService_Search_Call{Call: _e.mock.On("Search", _a0)}
}

func (_c *MockMediaService_Search_Call) Run(run func(_a0 *entity.Filter)) *MockMediaService_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Filter))
	})
	return _c
}

func (_c *MockMediaService_Search_Call) Return(_a0 *[]entity.MediaResult, _a1 error) *MockMediaService_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMediaService_Search_Call) RunAndReturn(run func(*entity.Filter) (*[]entity.MediaResult, error)) *MockMediaService_Search_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMediaService creates a new instance of MockMediaService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMediaService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMediaService {
	mock := &MockMediaService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
