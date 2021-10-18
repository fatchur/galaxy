package main

import (
	"log"

	"github.com/fatchur/galaxy.git/domain/translating"
	"github.com/fatchur/galaxy.git/domain/translating/models"
)

func main() {
	//translating.Translate("pish pish Iron is 3910 Credits")
	myRomanNum := models.BasicRomanNumber{}
	myRomanNum.Init()
	log.Println(">> basic number: ", myRomanNum)

	inputText := `glob is I
				prok is V
				pish is X
				tegj is L
				glob glob Silver is 34 Credits
				glob prok Gold is 57800 Credits
				pish pish Iron is 3910 Credits`
	translating.Translate(inputText, &myRomanNum)
	/*
		theTrader := trading.Trader{
			TraderID: 111,
			Word:     "glob prok Silver",
		}
		log.Println(theTrader)
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
	*/
}
