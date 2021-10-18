package usecase

import (
	"log"
	"strconv"

	"github.com/fatchur/galaxy.git/domain/translating/models"
)

func stringInSlice(list []string, a string) bool {
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

// InsertNumVal is ...
func InsertNumVal(word []string, alias *models.RomanNumberAlias) {
	alias.Alias = make(map[string]string)

	if stringInSlice(word, "I") {
		alias.Alias["I"] = word[0]
	} else if stringInSlice(word, "V") {
		alias.Alias["V"] = word[0]
	} else if stringInSlice(word, "X") {
		alias.Alias["X"] = word[0]
	} else if stringInSlice(word, "L") {
		alias.Alias["L"] = word[0]
	} else if stringInSlice(word, "C") {
		alias.Alias["C"] = word[0]
	} else if stringInSlice(word, "D") {
		alias.Alias["D"] = word[0]
	} else if stringInSlice(word, "M") {
		alias.Alias["M"] = word[0]
	}
}

// InsertCommodityVal is ...
func InsertCommodityVal(word []string) {
	indexOfIs, found := getItemIndex(word, "is")
	if found {
		commodityName := word[indexOfIs-1]
		log.Println(commodityName)
		price, _ := strconv.Atoi(word[indexOfIs+1])
		//myCommodity.Commodity[commodityName] =

	} else {
		log.Println("unknown pattern")
	}
}
