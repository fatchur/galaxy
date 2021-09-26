package usecase

import (
	"log"
	"strings"

	"github.com/fatchur/galaxy.git/domain/trading/models"
)

// Usecase is ...
type Usecase struct {
	SilverVal int
	GoldVal   int
	IronVal   float32
}

// Calculate is ...
func (u Usecase) Calculate(romanNumber *models.RomanNumberInterface, word string) {
	s := strings.Split(word, " ")

	totalVal := 0
	prevVal := 0
	for i, val := range s {
		log.Println(i, val)
		if val == "glob" {
			tmp := romanNumber
			totalVal, prevVal = romanNumber.Logic(prevVal, tmp, totalVal)
		}
		if val == "prok" {
			tmp := romanNumber.Prok
			totalVal, prevVal = Logic(prevVal, tmp, totalVal)
		}
		if val == "pish" {
			tmp := romanNumber.Pish
			totalVal, prevVal = Logic(prevVal, tmp, totalVal)
		}
		if val == "tegj" {
			tmp := romanNumber.Tegj
			totalVal, prevVal = Logic(prevVal, tmp, totalVal)
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
