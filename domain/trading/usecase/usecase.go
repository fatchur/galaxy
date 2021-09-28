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
	strangeWord := false
	for _, val := range s {
		totalVal, prevVal, strangeWord = myNumber.Logic(prevVal, totalVal, val)
		//log.Println(strangeWord, val)

		if strangeWord {
			if val == "Gold" || val == "gold" {
				totalVal = totalVal * u.GoldVal
			} else if val == "Silver" || val == "silver" {
				totalVal = totalVal * u.SilverVal
			} else if val == "Iron" || val == "iron" {
				totalVal = int(float32(totalVal) * u.IronVal)
			} else {
				totalVal = -1
				//log.Println(val)
				break
			}
		}
	}
	return totalVal
}
