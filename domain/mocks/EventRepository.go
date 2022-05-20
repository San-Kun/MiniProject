// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/San-Kun/MiniProject/domain"

	mock "github.com/stretchr/testify/mock"
)

// EventRepository is an autogenerated mock type for the EventRepository type
type EventRepository struct {
	mock.Mock
}

// CheckLogin provides a mock function with given fields: event
func (_m *EventRepository) CheckLogin(event *domain.Event) (*domain.Event, bool, error) {
	ret := _m.Called(event)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(*domain.Event) *domain.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*domain.Event) bool); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*domain.Event) error); ok {
		r2 = rf(event)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}


// Create provides a mock function with given fields: event
func (_m *EventRepository) Create(event *domain.Event) (*domain.Event, error) {
	ret := _m.Called(event)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(*domain.Event) *domain.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields:
func (_m *EventRepository) ReadAll() (*domain.Events, error) {
	ret := _m.Called()

	var r0 *domain.Events
	if rf, ok := ret.Get(0).(func() *domain.Events); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Events)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadByID provides a mock function with given fields: id
func (_m *EventRepository) ReadByID(id int) (*domain.Event, error) {
	ret := _m.Called(id)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(int) *domain.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: id
func (_m *EventRepository) Delete(id int) (*domain.Event, error) {
	ret := _m.Called(id)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(int) *domain.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
// UpdateByID provides a mock function with given fields: id
func (_m *EventRepository) Updates(id int) (*domain.Event, error) {
	ret := _m.Called(id)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(int) *domain.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}