// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]entities.User, error) {
	ret := _m.Called()

	var r0 []entities.User
	if rf, ok := ret.Get(0).(func() []entities.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.User)
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
