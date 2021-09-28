package usecase

import (
	"strings"

	"github.com/fatchur/galaxy.git/domain/trading/models"
)

// Usecase is ...
type Usecase struct {
	SilverVal int
	GoldVal   int
	IronVal   float32
}

// UsecaseInterface is ...
type UsecaseInterface interface {
	Calculate(romanNumber models.RomanNumberInterface, word string) int
}

// Calculate is ...
func (u Usecase) Calculate(myNumber models.RomanNumberInterface, word string) int {
	s := strings.Split(word, " ")

	totalVal := 0
	prevVal := 0
	for _, val := range s {
		totalVal, prevVal = myNumber.Logic(prevVal, totalVal, val)

		if val == "Gold" {
			totalVal = totalVal * u.GoldVal
		}
		if val == "Silver" {
			totalVal = totalVal * u.SilverVal
		}
		if val == "Iron" {
			totalVal = int(float32(totalVal) * u.IronVal)
		}
	}
	return totalVal
}
