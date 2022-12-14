// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package currency

import (
	"context"
	"sync"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/shopspring/decimal"
)

// Ensure, that CacheMock does implement Cache.
// If this is not the case, regenerate this file with moq.
var _ Cache = &CacheMock{}

// CacheMock is a mock implementation of Cache.
//
// 	func TestSomethingThatUsesCache(t *testing.T) {
//
// 		// make and configure a mocked Cache
// 		mockedCache := &CacheMock{
// 			GetRateFunc: func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
// 				panic("mock out the GetRate method")
// 			},
// 			SetRateFunc: func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
// 				panic("mock out the SetRate method")
// 			},
// 		}
//
// 		// use mockedCache in code that requires Cache
// 		// and then make assertions.
//
// 	}
type CacheMock struct {
	// GetRateFunc mocks the GetRate method.
	GetRateFunc func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error)

	// SetRateFunc mocks the SetRate method.
	SetRateFunc func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error

	// calls tracks calls to the methods.
	calls struct {
		// GetRate holds details about calls to the GetRate method.
		GetRate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Code is the code argument value.
			Code vos.CurrencyCode
		}
		// SetRate holds details about calls to the SetRate method.
		SetRate []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// CurrencyRates is the currencyRates argument value.
			CurrencyRates map[vos.CurrencyCode]decimal.Decimal
		}
	}
	lockGetRate sync.RWMutex
	lockSetRate sync.RWMutex
}

// GetRate calls GetRateFunc.
func (mock *CacheMock) GetRate(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
	if mock.GetRateFunc == nil {
		panic("CacheMock.GetRateFunc: method is nil but Cache.GetRate was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Code vos.CurrencyCode
	}{
		Ctx:  ctx,
		Code: code,
	}
	mock.lockGetRate.Lock()
	mock.calls.GetRate = append(mock.calls.GetRate, callInfo)
	mock.lockGetRate.Unlock()
	return mock.GetRateFunc(ctx, code)
}

// GetRateCalls gets all the calls that were made to GetRate.
// Check the length with:
//     len(mockedCache.GetRateCalls())
func (mock *CacheMock) GetRateCalls() []struct {
	Ctx  context.Context
	Code vos.CurrencyCode
} {
	var calls []struct {
		Ctx  context.Context
		Code vos.CurrencyCode
	}
	mock.lockGetRate.RLock()
	calls = mock.calls.GetRate
	mock.lockGetRate.RUnlock()
	return calls
}

// SetRate calls SetRateFunc.
func (mock *CacheMock) SetRate(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
	if mock.SetRateFunc == nil {
		panic("CacheMock.SetRateFunc: method is nil but Cache.SetRate was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		CurrencyRates   map[vos.CurrencyCode]decimal.Decimal
	}{
		ContextMoqParam: contextMoqParam,
		CurrencyRates:   currencyRates,
	}
	mock.lockSetRate.Lock()
	mock.calls.SetRate = append(mock.calls.SetRate, callInfo)
	mock.lockSetRate.Unlock()
	return mock.SetRateFunc(contextMoqParam, currencyRates)
}

// SetRateCalls gets all the calls that were made to SetRate.
// Check the length with:
//     len(mockedCache.SetRateCalls())
func (mock *CacheMock) SetRateCalls() []struct {
	ContextMoqParam context.Context
	CurrencyRates   map[vos.CurrencyCode]decimal.Decimal
} {
	var calls []struct {
		ContextMoqParam context.Context
		CurrencyRates   map[vos.CurrencyCode]decimal.Decimal
	}
	mock.lockSetRate.RLock()
	calls = mock.calls.SetRate
	mock.lockSetRate.RUnlock()
	return calls
}
