package galaxy

import (
	"log"
	"strings"

	"github.com/fatchur/galaxy.git/domain/trading/models"
	"github.com/fatchur/galaxy.git/domain/trading/usecase"
)

// UsecaseInterface is ...
type UsecaseInterface interface {
	Calculate(romanNumber *models.RomanNumber, word string)
}

// Calculate is ...
func (u usecase.Usecase) Calculate(romanNumber *models.RomanNumber, word string) {
	s := strings.Split(word, " ")

	totalVal := 0
	prevVal := 0
	for i, val := range s {
		log.Println(i, val)
		if val == "glob" {
			tmp := romanNumber.Glob
			totalVal, prevVal = usecase.Logic(prevVal, tmp, totalVal)
		}
		if val == "prok" {
			tmp := romanNumber.Prok
			totalVal, prevVal = usecase.Logic(prevVal, tmp, totalVal)
		}
		if val == "pish" {
			tmp := romanNumber.Pish
			totalVal, prevVal = usecase.Logic(prevVal, tmp, totalVal)
		}
		if val == "tegj" {
			tmp := romanNumber.Tegj
			totalVal, prevVal = usecase.Logic(prevVal, tmp, totalVal)
		}
		if val == "Gold" {
			totalVal = totalVal * u.GoldVal
		}
		if val == "Silver" {
			totalVal = totalVal * u.SilverVal
		}
		if val == "Iron" {
			totalVal = totalVal * int(u.IronVal)
		}
	}

	log.Println("==>", totalVal)

}
