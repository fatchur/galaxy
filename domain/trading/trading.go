package trading

import (
	"github.com/fatchur/galaxy.git/domain/trading/models"
	"github.com/fatchur/galaxy.git/domain/trading/usecase"
)

// Trader is ...
type Trader struct {
	Word     string
	TraderID int
}

// TraderInterface is ...
type TraderInterface interface {
	GetWord() string
}

// GetWord is ...
func (t Trader) GetWord() string {
	return t.Word
}

// DoTrading is ...
func DoTrading(myTrading TraderInterface, myUsecase usecase.UsecaseInterface, myNumber models.RomanNumberInterface) int {
	totalPrice := myUsecase.Calculate(myNumber, myTrading.GetWord())
	return totalPrice
}
