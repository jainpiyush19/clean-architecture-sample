// Code generated by mockery v1.0.0. DO NOT EDIT.

package adapters

import (
	context "context"

	wallet "github.com/jainpiyush19/cryptowallet/internal/domain/wallet"
	mock "github.com/stretchr/testify/mock"
)

// MockWalletRepository is an autogenerated mock type for the WalletRepository type
type MockWalletRepository struct {
	mock.Mock
}

// AllWallets provides a mock function with given fields: ctx
func (_m *MockWalletRepository) AllWallets(ctx context.Context) ([]*wallet.Info, error) {
	ret := _m.Called(ctx)

	var r0 []*wallet.Info
	if rf, ok := ret.Get(0).(func(context.Context) []*wallet.Info); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*wallet.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Balance provides a mock function with given fields: ctx, userID
func (_m *MockWalletRepository) Balance(ctx context.Context, userID int64) (*wallet.Info, error) {
	ret := _m.Called(ctx, userID)

	var r0 *wallet.Info
	if rf, ok := ret.Get(0).(func(context.Context, int64) *wallet.Info); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deposit provides a mock function with given fields: ctx, userID, amount
func (_m *MockWalletRepository) Deposit(ctx context.Context, userID int64, amount float64) error {
	ret := _m.Called(ctx, userID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, float64) error); ok {
		r0 = rf(ctx, userID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transfer provides a mock function with given fields: ctx, senderID, receiverID, amount
func (_m *MockWalletRepository) Transfer(ctx context.Context, senderID int64, receiverID int64, amount float64) error {
	ret := _m.Called(ctx, senderID, receiverID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, float64) error); ok {
		r0 = rf(ctx, senderID, receiverID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Withdraw provides a mock function with given fields: ctx, userID, amount
func (_m *MockWalletRepository) Withdraw(ctx context.Context, userID int64, amount float64) error {
	ret := _m.Called(ctx, userID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, float64) error); ok {
		r0 = rf(ctx, userID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
