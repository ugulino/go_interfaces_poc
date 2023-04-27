// Package mock_venda is a generated GoMock package.
package mock_venda

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

)

// MockCalculosVenda is a mock of CalculosVenda interface.
type MockCalculosVenda struct {
	ctrl     *gomock.Controller
	recorder *MockCalculosVendaMockRecorder
}

// MockCalculosVendaMockRecorder is the mock recorder for MockCalculosVenda.
type MockCalculosVendaMockRecorder struct {
	mock *MockCalculosVenda
}

// NewMockCalculosVenda creates a new mock instance.
func NewMockCalculosVenda(ctrl *gomock.Controller) *MockCalculosVenda {
	mock := &MockCalculosVenda{ctrl: ctrl}
	mock.recorder = &MockCalculosVendaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculosVenda) EXPECT() *MockCalculosVendaMockRecorder {
	return m.recorder
}

// CalculaComissao mocks base method.
func (m *MockCalculosVenda) CalculaComissao() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculaComissao")
	ret0 := ret[0].(float64)
        return ret0, nil
}

// CalculaComissao indicates an expected call of CalculaComissao.
func (mr *MockCalculosVendaMockRecorder) CalculaComissao() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculaComissao", reflect.TypeOf((*MockCalculosVenda)(nil).CalculaComissao))
}
