package models

import (
	"log"
	"testing"
)

type mockForTest struct {
	PrevVal   int
	InputWord string
	Output    int
	TestName  string
}

func TestLogic(t *testing.T) {
	myNumberMock := RomanNumber{
		Glob: 1,
		Prok: 5,
		Pish: 10,
		Tegj: 50,
	}

	var tableTest []mockForTest

	// glob
	// expected result 1
	myMock := mockForTest{
		PrevVal:   0,
		InputWord: "glob",
		Output:    1,
		TestName:  "case1",
	}
	tableTest = append(tableTest, myMock)

	// glob pish
	// expected result 9
	myMock = mockForTest{
		PrevVal:   1,
		InputWord: "pish",
		Output:    9,
		TestName:  "case2",
	}
	tableTest = append(tableTest, myMock)

	// glob tegj
	// expected result 49
	myMock = mockForTest{
		PrevVal:   1,
		InputWord: "tegj",
		Output:    49,
		TestName:  "case3",
	}
	tableTest = append(tableTest, myMock)

	// pish pish
	// expected result 20
	myMock = mockForTest{
		PrevVal:   10,
		InputWord: "pish",
		Output:    20,
		TestName:  "case4",
	}
	tableTest = append(tableTest, myMock)

	for idx, i := range tableTest {
		log.Println(idx, i.TestName)
		t.Run(i.TestName, func(t *testing.T) {
			actual, _, _ := myNumberMock.Logic(i.PrevVal, i.PrevVal, i.InputWord)

			if actual != i.Output {
				t.Error("error output", actual, i.Output)
			} else {
				t.Log("[PASS] ->>", i.Output, actual)
			}
		})
	}

}
