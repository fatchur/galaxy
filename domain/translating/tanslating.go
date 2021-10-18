package translating

import (
	"bufio"
	"log"
	"strings"

	"github.com/fatchur/galaxy.git/domain/translating/models"
	"github.com/fatchur/galaxy.git/domain/translating/usecase"
)

// Translate is ...
func Translate(input string, myRomanNum *models.BasicRomanNumber) {
	alias := models.RomanNumberAlias{}
	alias.Alias = make(map[string]string)
	commodity := models.Commodity{}
	commodity.Commodity = make(map[string]float32)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tmp := scanner.Text()
		tmp = strings.TrimSpace(tmp)
		splittedWord := usecase.Split(tmp)

		wordType := usecase.RecogWord(splittedWord)
		if wordType.NumberInit == true {
			usecase.InsertNumVal(splittedWord, &alias)
		} else if wordType.CommodityInit == true {
			usecase.InsertCommodityVal(splittedWord, &alias, myRomanNum, &commodity)
		}

		log.Println("==> alias: ", alias)
		log.Println("==> commodity: ", commodity)

	}
	/*
		//log.Println(wordType)
		log.Println(">>", splittedWord[0])
		log.Println(wordType)*/
}
