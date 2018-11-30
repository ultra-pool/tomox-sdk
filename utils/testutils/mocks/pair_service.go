// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bson "gopkg.in/mgo.v2/bson"
import common "github.com/ethereum/go-ethereum/common"
import mock "github.com/stretchr/testify/mock"

import types "github.com/tomochain/backend-matching-engine/types"

// PairService is an autogenerated mock type for the PairService type
type PairService struct {
	mock.Mock
}

// Create provides a mock function with given fields: pair
func (_m *PairService) Create(pair *types.Pair) error {
	ret := _m.Called(pair)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Pair) error); ok {
		r0 = rf(pair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *PairService) GetAll() ([]*types.Pair, error) {
	ret := _m.Called()

	var r0 []*types.Pair
	if rf, ok := ret.Get(0).(func() []*types.Pair); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Pair)
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

// GetByID provides a mock function with given fields: id
func (_m *PairService) GetByID(id bson.ObjectId) (*types.Pair, error) {
	ret := _m.Called(id)

	var r0 *types.Pair
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Pair); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pair)
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

// GetByTokenAddress provides a mock function with given fields: bt, qt
func (_m *PairService) GetByTokenAddress(bt common.Address, qt common.Address) (*types.Pair, error) {
	ret := _m.Called(bt, qt)

	var r0 *types.Pair
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) *types.Pair); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(bt, qt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PairService) GetTokenPairData(bt, qt common.Address) ([]*types.Tick, error) {
	ret := _m.Called(bt, qt)
	var r0 []*types.Tick
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) []*types.Tick); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Tick)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(bt, qt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PairService) GetAllTokenPairData() ([]*types.Tick, error) {
	ret := _m.Called()

	var r0 []*types.Tick
	if rf, ok := ret.Get(0).(func() []*types.Tick); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Tick)
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
