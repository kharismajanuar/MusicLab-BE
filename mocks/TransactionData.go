// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	transactions "musiclab-be/features/transactions"

	mock "github.com/stretchr/testify/mock"
)

// TransactionData is an autogenerated mock type for the TransactionData type
type TransactionData struct {
	mock.Mock
}

// GetMentorTransaction provides a mock function with given fields: mentorID, limit, offset
func (_m *TransactionData) GetMentorTransaction(mentorID uint, limit int, offset int) ([]transactions.Core, error) {
	ret := _m.Called(mentorID, limit, offset)

	var r0 []transactions.Core
	if rf, ok := ret.Get(0).(func(uint, int, int) []transactions.Core); ok {
		r0 = rf(mentorID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(mentorID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStudentTransaction provides a mock function with given fields: studentID, limit, offset
func (_m *TransactionData) GetStudentTransaction(studentID uint, limit int, offset int) ([]transactions.Core, error) {
	ret := _m.Called(studentID, limit, offset)

	var r0 []transactions.Core
	if rf, ok := ret.Get(0).(func(uint, int, int) []transactions.Core); ok {
		r0 = rf(studentID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(studentID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeTransaction provides a mock function with given fields: newTransaction
func (_m *TransactionData) MakeTransaction(newTransaction transactions.Core) error {
	ret := _m.Called(newTransaction)

	var r0 error
	if rf, ok := ret.Get(0).(func(transactions.Core) error); ok {
		r0 = rf(newTransaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectOne provides a mock function with given fields: transactionID
func (_m *TransactionData) SelectOne(transactionID uint) (transactions.Core, error) {
	ret := _m.Called(transactionID)

	var r0 transactions.Core
	if rf, ok := ret.Get(0).(func(uint) transactions.Core); ok {
		r0 = rf(transactionID)
	} else {
		r0 = ret.Get(0).(transactions.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: input
func (_m *TransactionData) UpdateTransaction(input transactions.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(transactions.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransactionData interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionData creates a new instance of TransactionData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionData(t mockConstructorTestingTNewTransactionData) *TransactionData {
	mock := &TransactionData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
