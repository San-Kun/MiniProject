// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/San-Kun/MiniProject/domain"

	mock "github.com/stretchr/testify/mock"
)

// PaymentTypeRepository is an autogenerated mock type for the PaymentTypeRepository type 
type PaymentTypeRepository struct {
	mock.Mock
}

// CheckLogin provides a mock function with given fields: PaymentType 
func (_m *PaymentTypeRepository) CheckLogin(paymentType *domain.PaymentType) (*domain.PaymentType, bool, error) {
	ret := _m.Called(paymentType)

	var r0 *domain.PaymentType
	if rf, ok := ret.Get(0).(func(*domain.PaymentType) *domain.PaymentType); ok {
		r0 = rf(paymentType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*domain.PaymentType) bool); ok {
		r1 = rf(paymentType)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*domain.PaymentType) error); ok {
		r2 = rf(paymentType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}


// Create provides a mock function with given fields: PaymentType
func (_m *PaymentTypeRepository) Create(paymentType *domain.PaymentType) (*domain.PaymentType, error) {
	ret := _m.Called(paymentType)

	var r0 *domain.PaymentType
	if rf, ok := ret.Get(0).(func(*domain.PaymentType) *domain.PaymentType); ok {
		r0 = rf(paymentType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.PaymentType) error); ok {
		r1 = rf(paymentType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields:
func (_m *PaymentTypeRepository) ReadAll() (*domain.PaymentTypes, error) {
	ret := _m.Called()

	var r0 *domain.PaymentTypes
	if rf, ok := ret.Get(0).(func() *domain.PaymentTypes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentTypes)
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
func (_m *PaymentTypeRepository) ReadByID(id int) (*domain.PaymentType, error) {
	ret := _m.Called(id)

	var r0 *domain.PaymentType
	if rf, ok := ret.Get(0).(func(int) *domain.PaymentType); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
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
func (_m *PaymentTypeRepository) Delete(id int) (*domain.PaymentType, error) {
	ret := _m.Called(id)

	var r0 *domain.PaymentType
	if rf, ok := ret.Get(0).(func(int) *domain.PaymentType); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
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
func (_m *PaymentTypeRepository) Updates(id int) (*domain.PaymentType, error) {
	ret := _m.Called(id)

	var r0 *domain.PaymentType
	if rf, ok := ret.Get(0).(func(int) *domain.PaymentType); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
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