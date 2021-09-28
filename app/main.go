package main

import (
	"log"

	"github.com/fatchur/galaxy.git/domain/trading"
	"github.com/fatchur/galaxy.git/domain/trading/models"
	"github.com/fatchur/galaxy.git/domain/trading/usecase"
)

func main() {
	theTrader := trading.Trader{
		TraderID: 111,
		Word:     "pish tegj glob glob",
	}

	price := usecase.Usecase{
		SilverVal: 17,
		GoldVal:   14450,
		IronVal:   196,
	}

	romanNum := models.RomanNumber{
		Glob: 1,
		Prok: 5,
		Pish: 10,
		Tegj: 50,
	}

	totalPrice := trading.DoTrading(theTrader, price, romanNum)
	log.Println(totalPrice)
}
