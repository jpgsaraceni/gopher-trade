// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package currency

import (
	"context"
	"sync"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			GetCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
// 				panic("mock out the GetCurrencyByCode method")
// 			},
// 			UpsertCurrencyFunc: func(ctx context.Context, cur entities.Currency) (entities.Currency, error) {
// 				panic("mock out the UpsertCurrency method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// GetCurrencyByCodeFunc mocks the GetCurrencyByCode method.
	GetCurrencyByCodeFunc func(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error)

	// UpsertCurrencyFunc mocks the UpsertCurrency method.
	UpsertCurrencyFunc func(ctx context.Context, cur entities.Currency) (entities.Currency, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetCurrencyByCode holds details about calls to the GetCurrencyByCode method.
		GetCurrencyByCode []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Code is the code argument value.
			Code vos.CurrencyCode
		}
		// UpsertCurrency holds details about calls to the UpsertCurrency method.
		UpsertCurrency []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Cur is the cur argument value.
			Cur entities.Currency
		}
	}
	lockGetCurrencyByCode sync.RWMutex
	lockUpsertCurrency    sync.RWMutex
}

// GetCurrencyByCode calls GetCurrencyByCodeFunc.
func (mock *RepositoryMock) GetCurrencyByCode(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
	if mock.GetCurrencyByCodeFunc == nil {
		panic("RepositoryMock.GetCurrencyByCodeFunc: method is nil but Repository.GetCurrencyByCode was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Code vos.CurrencyCode
	}{
		Ctx:  ctx,
		Code: code,
	}
	mock.lockGetCurrencyByCode.Lock()
	mock.calls.GetCurrencyByCode = append(mock.calls.GetCurrencyByCode, callInfo)
	mock.lockGetCurrencyByCode.Unlock()
	return mock.GetCurrencyByCodeFunc(ctx, code)
}

// GetCurrencyByCodeCalls gets all the calls that were made to GetCurrencyByCode.
// Check the length with:
//     len(mockedRepository.GetCurrencyByCodeCalls())
func (mock *RepositoryMock) GetCurrencyByCodeCalls() []struct {
	Ctx  context.Context
	Code vos.CurrencyCode
} {
	var calls []struct {
		Ctx  context.Context
		Code vos.CurrencyCode
	}
	mock.lockGetCurrencyByCode.RLock()
	calls = mock.calls.GetCurrencyByCode
	mock.lockGetCurrencyByCode.RUnlock()
	return calls
}

// UpsertCurrency calls UpsertCurrencyFunc.
func (mock *RepositoryMock) UpsertCurrency(ctx context.Context, cur entities.Currency) (entities.Currency, error) {
	if mock.UpsertCurrencyFunc == nil {
		panic("RepositoryMock.UpsertCurrencyFunc: method is nil but Repository.UpsertCurrency was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Cur entities.Currency
	}{
		Ctx: ctx,
		Cur: cur,
	}
	mock.lockUpsertCurrency.Lock()
	mock.calls.UpsertCurrency = append(mock.calls.UpsertCurrency, callInfo)
	mock.lockUpsertCurrency.Unlock()
	return mock.UpsertCurrencyFunc(ctx, cur)
}

// UpsertCurrencyCalls gets all the calls that were made to UpsertCurrency.
// Check the length with:
//     len(mockedRepository.UpsertCurrencyCalls())
func (mock *RepositoryMock) UpsertCurrencyCalls() []struct {
	Ctx context.Context
	Cur entities.Currency
} {
	var calls []struct {
		Ctx context.Context
		Cur entities.Currency
	}
	mock.lockUpsertCurrency.RLock()
	calls = mock.calls.UpsertCurrency
	mock.lockUpsertCurrency.RUnlock()
	return calls
}
