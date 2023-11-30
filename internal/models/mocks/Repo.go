// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	todo "github.com/vlasashk/todo-manager/internal/models/todo"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *Repo) CreateTask(task todo.TaskReq) (todo.Task, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 todo.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(todo.TaskReq) (todo.Task, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(todo.TaskReq) todo.Task); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(todo.Task)
	}

	if rf, ok := ret.Get(1).(func(todo.TaskReq) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: taskID
func (_m *Repo) DeleteTask(taskID string) error {
	ret := _m.Called(taskID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTask provides a mock function with given fields: taskID
func (_m *Repo) GetTask(taskID string) (todo.Task, error) {
	ret := _m.Called(taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetTask")
	}

	var r0 todo.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (todo.Task, error)); ok {
		return rf(taskID)
	}
	if rf, ok := ret.Get(0).(func(string) todo.Task); ok {
		r0 = rf(taskID)
	} else {
		r0 = ret.Get(0).(todo.Task)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: page, date, status
func (_m *Repo) ListTasks(page uint, date string, status string) ([]todo.Task, error) {
	ret := _m.Called(page, date, status)

	if len(ret) == 0 {
		panic("no return value specified for ListTasks")
	}

	var r0 []todo.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, string) ([]todo.Task, error)); ok {
		return rf(page, date, status)
	}
	if rf, ok := ret.Get(0).(func(uint, string, string) []todo.Task); ok {
		r0 = rf(page, date, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]todo.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, string) error); ok {
		r1 = rf(page, date, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: task, taskID
func (_m *Repo) UpdateTask(task todo.TaskReq, taskID string) (todo.Task, error) {
	ret := _m.Called(task, taskID)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 todo.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(todo.TaskReq, string) (todo.Task, error)); ok {
		return rf(task, taskID)
	}
	if rf, ok := ret.Get(0).(func(todo.TaskReq, string) todo.Task); ok {
		r0 = rf(task, taskID)
	} else {
		r0 = ret.Get(0).(todo.Task)
	}

	if rf, ok := ret.Get(1).(func(todo.TaskReq, string) error); ok {
		r1 = rf(task, taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
