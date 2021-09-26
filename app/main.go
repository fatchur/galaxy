package main

import (
	"log"

	"github.com/fatchur/galaxy.git/domain/trading/models"
	"github.com/fatchur/galaxy.git/domain/trading/usecase"
)

func main() {
	price := usecase.Usecase{
		SilverVal: 17,
		GoldVal:   14450,
		IronVal:   195.5,
	}

	romanNum := models.RomanNumber{
		Glob: 1,
		Prok: 5,
		Pish: 10,
		Tegj: 50,
	}

	price.Calculate(&romanNum, "glob prok Iron")
	log.Println(price)
}
