// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/iki-rumondor/project1_grup9/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *TaskRepository) Delete(_a0 *domain.Task) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Task) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *TaskRepository) FindAll() ([]domain.Task, error) {
	ret := _m.Called()

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: _a0
func (_m *TaskRepository) FindByID(_a0 uint) (*domain.Task, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.Task, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: _a0
func (_m *TaskRepository) Upsert(_a0 *domain.Task) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Task) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}