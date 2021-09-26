package models

// RomanNumber is ...
type RomanNumber struct {
	Glob int
	Prok int
	Pish int
	Tegj int
}

// Logic is ...
func (r RomanNumber) Logic(prevVal int, currentVal int, totalVal int) (int, int) {
	if prevVal < currentVal && prevVal != 0 {
		totalVal = totalVal - prevVal + (currentVal - prevVal)
	} else {
		totalVal += currentVal
	}
	prevVal = currentVal
	return totalVal, prevVal
}

// RomanNumberInterface is ...
type RomanNumberInterface interface {
	Logic(prevVal int, currentVal int, totalVal int) (int, int)
}
