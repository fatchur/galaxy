package usecase

import (
	"sort"
)

// WordType is ...
type WordType struct {
	NumberInit    bool
	CommodityInit bool
	Question      bool
}

// WordTypeInterface is ...
type WordTypeInterface interface {
}

//RecogWord is ...
func RecogWord(mySlice []string) WordType {

	myWordType := WordType{false, false, false}

	if len(mySlice) == 3 {
		myWordType.NumberInit = true
	} else if int(sort.SearchStrings(mySlice, "Credits")) >= 1 {
		myWordType.CommodityInit = true
	} else if int(sort.SearchStrings(mySlice, "?")) >= 1 {
		myWordType.Question = true
	}

	return myWordType
}
