// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bson "github.com/globalsign/mgo/bson"
import common "github.com/ethereum/go-ethereum/common"
import mock "github.com/stretchr/testify/mock"

import types "github.com/tomochain/tomox-sdk/types"

// AccountService is an autogenerated mock type for the AccountService type
type AccountService struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *AccountService) Create(account *types.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *AccountService) GetAll() ([]types.Account, error) {
	ret := _m.Called()

	var r0 []types.Account
	if rf, ok := ret.Get(0).(func() []types.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Account)
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

// GetByAddress provides a mock function with given fields: a
func (_m *AccountService) GetByAddress(a common.Address) (*types.Account, error) {
	ret := _m.Called(a)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(common.Address) *types.Account); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *AccountService) GetByID(id bson.ObjectId) (*types.Account, error) {
	ret := _m.Called(id)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalance provides a mock function with given fields: owner, token
func (_m *AccountService) GetTokenBalance(owner common.Address, token common.Address) (*types.TokenBalance, error) {
	ret := _m.Called(owner, token)

	var r0 *types.TokenBalance
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) *types.TokenBalance); ok {
		r0 = rf(owner, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TokenBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(owner, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalances provides a mock function with given fields: owner
func (_m *AccountService) GetTokenBalances(owner common.Address) (map[common.Address]*types.TokenBalance, error) {
	ret := _m.Called(owner)

	var r0 map[common.Address]*types.TokenBalance
	if rf, ok := ret.Get(0).(func(common.Address) map[common.Address]*types.TokenBalance); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Address]*types.TokenBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
