// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	graphql "github.com/machinebox/graphql"

	mock "github.com/stretchr/testify/mock"
)

// GraphQLClient is an autogenerated mock type for the GraphQLClient type
type GraphQLClient struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, req, resp
func (_m *GraphQLClient) Run(ctx context.Context, req *graphql.Request, resp interface{}) error {
	ret := _m.Called(ctx, req, resp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *graphql.Request, interface{}) error); ok {
		r0 = rf(ctx, req, resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewGraphQLClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGraphQLClient creates a new instance of GraphQLClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGraphQLClient(t mockConstructorTestingTNewGraphQLClient) *GraphQLClient {
	mock := &GraphQLClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
