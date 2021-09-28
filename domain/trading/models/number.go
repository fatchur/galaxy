package models

// RomanNumber is ...
type RomanNumber struct {
	Glob int
	Prok int
	Pish int
	Tegj int
}

// RomanNumberInterface is ...
type RomanNumberInterface interface {
	Logic(prevVal int, currentVal int, totalVal string) (int, int)
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

// Logic is ...
func (r RomanNumber) Logic(prevVal int, totalVal int, word string) (int, int) {
	if word == "glob" {
		tmp := r.Glob
		totalVal, prevVal = util(prevVal, tmp, totalVal)
	} else if word == "prok" {
		tmp := r.Prok
		totalVal, prevVal = util(prevVal, tmp, totalVal)
	} else if word == "pish" {
		tmp := r.Pish
		totalVal, prevVal = util(prevVal, tmp, totalVal)
	} else if word == "tegj" {
		tmp := r.Tegj
		totalVal, prevVal = util(prevVal, tmp, totalVal)
	}

	return totalVal, prevVal
}
