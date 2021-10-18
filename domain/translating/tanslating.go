package translating

import (
	"log"

	"github.com/fatchur/galaxy.git/domain/translating/models"
	"github.com/fatchur/galaxy.git/domain/translating/usecase"
)

// Translate is ...
func Translate(myWord string) {
	splittedWord := usecase.Split(myWord)
	wordType := usecase.RecogWord(splittedWord)
	alias := models.RomanNumberAlias{}

	if wordType.NumberInit == true {
		usecase.InsertNumVal(splittedWord, &alias)
		log.Println(alias)
	}

	usecase.InsertCommodityVal(splittedWord)

	//log.Println(wordType)
	log.Println(">>", splittedWord[0])
	log.Println(wordType)
}
