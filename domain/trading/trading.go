package galaxy

import (
	"github.com/fatchur/galaxy.git/domain/trading/models"
)

// UsecaseInterface is ...
type UsecaseInterface interface {
	Calculate(romanNumber *models.RomanNumber, word string)
}
