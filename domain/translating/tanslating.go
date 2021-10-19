package translating

import (
	"bufio"
	"strings"

	"github.com/fatchur/galaxy.git/domain/translating/models"
	"github.com/fatchur/galaxy.git/domain/translating/usecase"
)

// Translate is ...
func Translate(input string, myRomanNum *models.BasicRomanNumber, alias *models.RomanNumberAlias, commodity *models.Commodity) [][]string {

	var questionList [][]string

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tmp := scanner.Text()
		tmp = strings.TrimSpace(tmp)
		splittedWord := usecase.Split(tmp)

		wordType := usecase.RecogWord(splittedWord)
		if wordType.NumberInit == true {
			usecase.InsertNumVal(splittedWord, alias)
		} else if wordType.CommodityInit == true {
			usecase.InsertCommodityVal(splittedWord, alias, myRomanNum, commodity)
		} else {
			questionList = append(questionList, splittedWord)
		}
	}
	return questionList
}
