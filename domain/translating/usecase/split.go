package usecase

import (
	"strings"
)

// Split is ...
func Split(myText string) []string {
	s := strings.Split(myText, " ")
	return s
}
