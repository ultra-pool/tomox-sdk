// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/ethereum/go-ethereum/common"

import mock "github.com/stretchr/testify/mock"
import types "github.com/tomochain/backend-matching-engine/types"
import ws "github.com/tomochain/backend-matching-engine/ws"

// TradeService is an autogenerated mock type for the TradeService type
type TradeService struct {
	mock.Mock
}

// GetByHash provides a mock function with given fields: hash
func (_m *TradeService) GetByHash(hash common.Hash) (*types.Trade, error) {
	ret := _m.Called(hash)

	var r0 *types.Trade
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Trade); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByOrderHash provides a mock function with given fields: hash
func (_m *TradeService) GetByOrderHash(hash common.Hash) ([]*types.Trade, error) {
	ret := _m.Called(hash)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Hash) []*types.Trade); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPairAddress provides a mock function with given fields: bt, qt
func (_m *TradeService) GetByPairAddress(bt common.Address, qt common.Address) ([]*types.Trade, error) {
	ret := _m.Called(bt, qt)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) []*types.Trade); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
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

// GetByPairName provides a mock function with given fields: p
func (_m *TradeService) GetByPairName(p string) ([]*types.Trade, error) {
	ret := _m.Called(p)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(string) []*types.Trade); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserAddress provides a mock function with given fields: addr
func (_m *TradeService) GetByUserAddress(addr common.Address) ([]*types.Trade, error) {
	ret := _m.Called(addr)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address) []*types.Trade); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TradeService) GetByOrderHashes(h []common.Hash) ([]*types.Trade, error) {
	ret := _m.Called(h)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func([]common.Hash) []*types.Trade); ok {
		r0 = rf(h)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]common.Hash) error); ok {
		r1 = rf(h)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrades provides a mock function with given fields: bt, qt
func (_m *TradeService) GetByMakerOrderHash(h common.Hash) ([]*types.Trade, error) {
	ret := _m.Called(h)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Hash) []*types.Trade); ok {
		r0 = rf(h)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(h)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TradeService) GetSortedTradesByUserAddress(a common.Address, limit ...int) ([]*types.Trade, error) {

	ret := _m.Called(a, limit)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address, ...int) []*types.Trade); ok {
		r0 = rf(a, limit...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, ...int) error); ok {
		r1 = rf(a, limit...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TradeService) GetSortedTrades(bt, qt common.Address, n int) ([]*types.Trade, error) {
	ret := _m.Called(bt, qt)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) []*types.Trade); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
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

func (_m *TradeService) GetByTakerOrderHash(h common.Hash) ([]*types.Trade, error) {
	ret := _m.Called(h)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Hash) []*types.Trade); ok {
		r0 = rf(h)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(h)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TradeService) GetAllTradesByPairAddress(bt, qt common.Address) ([]*types.Trade, error) {
	ret := _m.Called(bt, qt)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) []*types.Trade); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
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

// Subscribe provides a mock function with given fields: conn, bt, qt
func (_m *TradeService) Unsubscribe(c *ws.Client) {
	_m.Called(c)
}

// Unsubscribe provides a mock function with given fields: conn, bt, qt
func (_m *TradeService) UnsubscribeChannel(c *ws.Client, bt, qt common.Address) {
	_m.Called(c, bt, qt)
}

// UpdateTradeTxHash provides a mock function with given fields: tr, txHash
func (_m *TradeService) UpdateTradeTxHash(tr *types.Trade, txHash common.Hash) error {
	ret := _m.Called(tr, txHash)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Trade, common.Hash) error); ok {
		r0 = rf(tr, txHash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
