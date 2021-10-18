package usecase

import (
	"log"
	"reflect"
	"strconv"

	"github.com/fatchur/galaxy.git/domain/translating/models"
)

// StringInSlice is ...
func StringInSlice(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getItemIndex(list []string, a string) (int, bool) {
	selectedIndex := 0
	getItem := false

	for i, val := range list {
		if val == a {
			getItem = true
			selectedIndex = i
		}
	}
	return selectedIndex, getItem
}

func util(prevVal int, currentVal int, totalVal int) (int, int) {
	if prevVal < currentVal && prevVal != 0 {
		totalVal = totalVal - prevVal + (currentVal - prevVal)
	} else {
		totalVal += currentVal
	}
	prevVal = currentVal
	return totalVal, prevVal
}

// CalculateRomanValue is ...
func CalculateRomanValue(word []string, number *models.BasicRomanNumber, alias *models.RomanNumberAlias) int {
	totalVal := 0
	prevVal := 0
	//strangeWord := false

	for _, val := range word {
		currentValInRoman := alias.Alias[val]
		ref := reflect.ValueOf(number)
		currentValInNum := reflect.Indirect(ref).FieldByName(currentValInRoman).Interface().(int)
		totalVal, prevVal = util(prevVal, currentValInNum, totalVal)
	}
	return totalVal
}

// InsertNumVal is ...
func InsertNumVal(word []string, alias *models.RomanNumberAlias) {
	if StringInSlice(word, "I") {
		alias.Alias[word[0]] = "I"
	} else if StringInSlice(word, "V") {
		alias.Alias[word[0]] = "V"
	} else if StringInSlice(word, "X") {
		alias.Alias[word[0]] = "X"
	} else if StringInSlice(word, "L") {
		alias.Alias[word[0]] = "L"
	} else if StringInSlice(word, "C") {
		alias.Alias[word[0]] = "C"
	} else if StringInSlice(word, "D") {
		alias.Alias[word[0]] = "D"
	} else if StringInSlice(word, "M") {
		alias.Alias[word[0]] = "M"
	}
}

// InsertCommodityVal is ...
func InsertCommodityVal(word []string, alias *models.RomanNumberAlias,
	myNum *models.BasicRomanNumber,
	myCommodity *models.Commodity) {
	indexOfIs, found := getItemIndex(word, "is")
	if found {
		commodityName := word[indexOfIs-1]
		log.Println(commodityName, word[:indexOfIs-1])
		totalVal := CalculateRomanValue(word[:indexOfIs-1], myNum, alias)
		price, _ := strconv.Atoi(word[indexOfIs+1])
		myCommodity.Commodity[commodityName] = float32(price) / float32(totalVal)

	} else {
		log.Println("unknown pattern")
	}
}
