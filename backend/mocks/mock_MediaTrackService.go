// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "backend/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockMediaTrackService is an autogenerated mock type for the MediaTrackService type
type MockMediaTrackService struct {
	mock.Mock
}

type MockMediaTrackService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMediaTrackService) EXPECT() *MockMediaTrackService_Expecter {
	return &MockMediaTrackService_Expecter{mock: &_m.Mock}
}

// Change provides a mock function with given fields: _a0
func (_m *MockMediaTrackService) Change(_a0 *entity.MediaTrackBase) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Change")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.MediaTrackBase) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaTrackService_Change_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Change'
type MockMediaTrackService_Change_Call struct {
	*mock.Call
}

// Change is a helper method to define mock.On call
//   - _a0 *entity.MediaTrackBase
func (_e *MockMediaTrackService_Expecter) Change(_a0 interface{}) *MockMediaTrackService_Change_Call {
	return &MockMediaTrackService_Change_Call{Call: _e.mock.On("Change", _a0)}
}

func (_c *MockMediaTrackService_Change_Call) Run(run func(_a0 *entity.MediaTrackBase)) *MockMediaTrackService_Change_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.MediaTrackBase))
	})
	return _c
}

func (_c *MockMediaTrackService_Change_Call) Return(_a0 error) *MockMediaTrackService_Change_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaTrackService_Change_Call) RunAndReturn(run func(*entity.MediaTrackBase) error) *MockMediaTrackService_Change_Call {
	_c.Call.Return(run)
	return _c
}

// LoadAll provides a mock function with given fields: userID
func (_m *MockMediaTrackService) LoadAll(userID uint) (*[]entity.MediaTrackView, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for LoadAll")
	}

	var r0 *[]entity.MediaTrackView
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*[]entity.MediaTrackView, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) *[]entity.MediaTrackView); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.MediaTrackView)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMediaTrackService_LoadAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadAll'
type MockMediaTrackService_LoadAll_Call struct {
	*mock.Call
}

// LoadAll is a helper method to define mock.On call
//   - userID uint
func (_e *MockMediaTrackService_Expecter) LoadAll(userID interface{}) *MockMediaTrackService_LoadAll_Call {
	return &MockMediaTrackService_LoadAll_Call{Call: _e.mock.On("LoadAll", userID)}
}

func (_c *MockMediaTrackService_LoadAll_Call) Run(run func(userID uint)) *MockMediaTrackService_LoadAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockMediaTrackService_LoadAll_Call) Return(_a0 *[]entity.MediaTrackView, _a1 error) *MockMediaTrackService_LoadAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMediaTrackService_LoadAll_Call) RunAndReturn(run func(uint) (*[]entity.MediaTrackView, error)) *MockMediaTrackService_LoadAll_Call {
	_c.Call.Return(run)
	return _c
}

// Track provides a mock function with given fields: _a0
func (_m *MockMediaTrackService) Track(_a0 *entity.MediaTrackBase) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Track")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.MediaTrackBase) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaTrackService_Track_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Track'
type MockMediaTrackService_Track_Call struct {
	*mock.Call
}

// Track is a helper method to define mock.On call
//   - _a0 *entity.MediaTrackBase
func (_e *MockMediaTrackService_Expecter) Track(_a0 interface{}) *MockMediaTrackService_Track_Call {
	return &MockMediaTrackService_Track_Call{Call: _e.mock.On("Track", _a0)}
}

func (_c *MockMediaTrackService_Track_Call) Run(run func(_a0 *entity.MediaTrackBase)) *MockMediaTrackService_Track_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.MediaTrackBase))
	})
	return _c
}

func (_c *MockMediaTrackService_Track_Call) Return(_a0 error) *MockMediaTrackService_Track_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaTrackService_Track_Call) RunAndReturn(run func(*entity.MediaTrackBase) error) *MockMediaTrackService_Track_Call {
	_c.Call.Return(run)
	return _c
}

// Untrack provides a mock function with given fields: userID, mediaID
func (_m *MockMediaTrackService) Untrack(userID uint, mediaID uint) error {
	ret := _m.Called(userID, mediaID)

	if len(ret) == 0 {
		panic("no return value specified for Untrack")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, mediaID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaTrackService_Untrack_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Untrack'
type MockMediaTrackService_Untrack_Call struct {
	*mock.Call
}

// Untrack is a helper method to define mock.On call
//   - userID uint
//   - mediaID uint
func (_e *MockMediaTrackService_Expecter) Untrack(userID interface{}, mediaID interface{}) *MockMediaTrackService_Untrack_Call {
	return &MockMediaTrackService_Untrack_Call{Call: _e.mock.On("Untrack", userID, mediaID)}
}

func (_c *MockMediaTrackService_Untrack_Call) Run(run func(userID uint, mediaID uint)) *MockMediaTrackService_Untrack_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *MockMediaTrackService_Untrack_Call) Return(_a0 error) *MockMediaTrackService_Untrack_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaTrackService_Untrack_Call) RunAndReturn(run func(uint, uint) error) *MockMediaTrackService_Untrack_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMediaTrackService creates a new instance of MockMediaTrackService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMediaTrackService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMediaTrackService {
	mock := &MockMediaTrackService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
