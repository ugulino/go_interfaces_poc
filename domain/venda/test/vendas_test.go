package mock_venda

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	domain "varejo.org.br/domain/venda"
	mock "varejo.org.br/domain/venda/test/mock_venda"
)

type VendaUserCaseSuite struct {
	suite.Suite
	ctrl          *gomock.Controller
	Venda         *domain.Venda
	CalculosVenda *mock.MockCalculosVenda
}

func TestVendaUserCaseSuite(t *testing.T) {
	suite.Run(t, new(VendaUserCaseSuite))
}

func (s VendaUserCaseSuite) TestVendaUserCaseSuiteDown() {
	defer s.ctrl.Finish()
}

func (s *VendaUserCaseSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.CalculosVenda = mock.NewMockCalculosVenda(s.ctrl)

	s.CalculosVenda.EXPECT().CalculaComissao().Return(100.00, nil).AnyTimes()

	s.Venda = &domain.Venda{
		Valor:    1000,
		Status:   "Faturada",
		Calculos: s.CalculosVenda,
	}
}

func (s VendaUserCaseSuite) TestVendaUserCaseSuiteSuccess() {
	var comissao, _ = s.Venda.Execute(s.CalculosVenda)
	s.Assert().True(comissao == 100.00)
}

func (s VendaUserCaseSuite) TestVendaUserCaseSuiteCancelada() {
	s.CalculosVenda = mock.NewMockCalculosVenda(s.ctrl)

	s.Venda.Status = "Cancelada"
	s.CalculosVenda.EXPECT().CalculaComissao().Return(0.00, nil).AnyTimes()

	var comissao, _ = s.Venda.Execute(s.CalculosVenda)
	s.Assert().True(comissao == 0.00)
}
