// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/github"

	mock "github.com/stretchr/testify/mock"
)

// RepositoriesService is an autogenerated mock type for the RepositoriesService type
type RepositoriesService struct {
	mock.Mock
}

// ListStatuses provides a mock function with given fields: ctx, owner, repo, ref, opt
func (_m *RepositoriesService) ListStatuses(ctx context.Context, owner string, repo string, ref string, opt *github.ListOptions) ([]*github.RepoStatus, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, ref, opt)

	if len(ret) == 0 {
		panic("no return value specified for ListStatuses")
	}

	var r0 []*github.RepoStatus
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.ListOptions) ([]*github.RepoStatus, *github.Response, error)); ok {
		return rf(ctx, owner, repo, ref, opt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.ListOptions) []*github.RepoStatus); ok {
		r0 = rf(ctx, owner, repo, ref, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.RepoStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, ref, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, owner, repo, ref, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewRepositoriesService creates a new instance of RepositoriesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoriesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoriesService {
	mock := &RepositoriesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
