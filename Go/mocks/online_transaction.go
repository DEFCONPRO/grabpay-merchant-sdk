// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package gemocks

import (
	dto "github.com/grab/grabpay-merchant-sdk/dto"
	context "context"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// OnlineTransaction is an autogenerated mock type for the OnlineTransaction type
type OnlineTransaction struct {
	mock.Mock
}

// OnaChargeComplete provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaChargeComplete(ctx context.Context, params *dto.OnaChargeCompleteParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaChargeCompleteParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaChargeCompleteParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaChargeInit provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaChargeInit(ctx context.Context, params *dto.OnaChargeInitParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaChargeInitParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaChargeInitParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaCreateWebUrl provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaCreateWebUrl(ctx context.Context, params *dto.OnaCreateWebUrlParams) (string, error) {
	ret := _m.Called(ctx, params)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaCreateWebUrlParams) string); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaCreateWebUrlParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaGetChargeStatus provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaGetChargeStatus(ctx context.Context, params *dto.OnaGetChargeStatusParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaGetChargeStatusParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaGetChargeStatusParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaGetOTCStatus provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaGetOTCStatus(ctx context.Context, params *dto.OnaGetOTCStatusParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaGetOTCStatusParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaGetOTCStatusParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaGetRefundStatus provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaGetRefundStatus(ctx context.Context, params *dto.OnaGetRefundStatusParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaGetRefundStatusParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaGetRefundStatusParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaOAuth2Token provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaOAuth2Token(ctx context.Context, params *dto.OnaOAuth2TokenParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaOAuth2TokenParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaOAuth2TokenParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnaRefund provides a mock function with given fields: ctx, params
func (_m *OnlineTransaction) OnaRefund(ctx context.Context, params *dto.OnaRefundParams) (*http.Response, error) {
	ret := _m.Called(ctx, params)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *dto.OnaRefundParams) *http.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.OnaRefundParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
