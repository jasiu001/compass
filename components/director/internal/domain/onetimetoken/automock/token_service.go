// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// GenerateOneTimeToken provides a mock function with given fields: ctx, runtimeID, tokenType
func (_m *TokenService) GenerateOneTimeToken(ctx context.Context, runtimeID string, tokenType model.SystemAuthReferenceObjectType) (model.OneTimeToken, error) {
	ret := _m.Called(ctx, runtimeID, tokenType)

	var r0 model.OneTimeToken
	if rf, ok := ret.Get(0).(func(context.Context, string, model.SystemAuthReferenceObjectType) model.OneTimeToken); ok {
		r0 = rf(ctx, runtimeID, tokenType)
	} else {
		r0 = ret.Get(0).(model.OneTimeToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.SystemAuthReferenceObjectType) error); ok {
		r1 = rf(ctx, runtimeID, tokenType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
