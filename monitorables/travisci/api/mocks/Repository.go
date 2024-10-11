// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/monitorables/travisci/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetLastBuildStatus provides a mock function with given fields: owner, repository, branch
func (_m *Repository) GetLastBuildStatus(owner string, repository string, branch string) (*models.Build, error) {
	ret := _m.Called(owner, repository, branch)

	if len(ret) == 0 {
		panic("no return value specified for GetLastBuildStatus")
	}

	var r0 *models.Build
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*models.Build, error)); ok {
		return rf(owner, repository, branch)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Build); ok {
		r0 = rf(owner, repository, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Build)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repository, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
