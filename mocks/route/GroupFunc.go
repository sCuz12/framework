// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	route "github.com/goravel/framework/contracts/route"
	mock "github.com/stretchr/testify/mock"
)

// GroupFunc is an autogenerated mock type for the GroupFunc type
type GroupFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: router
func (_m *GroupFunc) Execute(router route.Router) {
	_m.Called(router)
}

// NewGroupFunc creates a new instance of GroupFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupFunc {
	mock := &GroupFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
