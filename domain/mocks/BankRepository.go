// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/San-Kun/MiniProject/domain"

	mock "github.com/stretchr/testify/mock"
)

// BankRepository is an autogenerated mock type for the BankRepository type
type BankRepository struct {
	mock.Mock
}

// CheckLogin provides a mock function with given fields: Bank
func (_m *BankRepository) CheckLogin(bank *domain.Bank) (*domain.Bank, bool, error) {
	ret := _m.Called(bank)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(*domain.Bank) *domain.Bank); ok {
		r0 = rf(bank)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*domain.Bank) bool); ok {
		r1 = rf(bank)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*domain.Bank) error); ok {
		r2 = rf(bank)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}


// Create provides a mock function with given fields: Bank
func (_m *BankRepository) Create(bank *domain.Bank) (*domain.Bank, error) {
	ret := _m.Called(bank)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(*domain.Bank) *domain.Bank); ok {
		r0 = rf(bank)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Bank) error); ok {
		r1 = rf(bank)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields:
func (_m *BankRepository) ReadAll() (*domain.Banks, error) {
	ret := _m.Called()

	var r0 *domain.Banks
	if rf, ok := ret.Get(0).(func() *domain.Banks); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Banks)
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
func (_m *BankRepository) ReadByID(id int) (*domain.Bank, error) {
	ret := _m.Called(id)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(int) *domain.Bank); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
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
func (_m *BankRepository) Delete(id int) (*domain.Bank, error) {
	ret := _m.Called(id)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(int) *domain.Bank); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
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
func (_m *BankRepository) Updates(id int) (*domain.Bank, error) {
	ret := _m.Called(id)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(int) *domain.Bank); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
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