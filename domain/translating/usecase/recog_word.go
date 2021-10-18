package usecase

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
	} else if StringInSlice(mySlice, "Credits") {
		myWordType.CommodityInit = true
	} else if StringInSlice(mySlice, "?") {
		myWordType.Question = true
	}

	return myWordType
}
