package usecase

import (
	"log"
	"testing"

	"github.com/fatchur/galaxy.git/domain/trading/models"
)

type mockForTest struct {
	Input    string
	Output   int
	TestName string
}

func TestCalculate(t *testing.T) {
	log.Println("=================")
	myNumberMock := models.RomanNumber{
		Glob: 1,
		Prok: 5,
		Pish: 10,
		Tegj: 50,
	}
	myUsecaseMock := Usecase{
		SilverVal: 17,
		GoldVal:   14450,
		IronVal:   195.5,
	}
	var tableTest []mockForTest

	// case 1, pish tegj glob glob
	// expected result 42
	myMock := mockForTest{
		Input:    "pish tegj glob glob",
		Output:   42,
		TestName: "case1",
	}
	tableTest = append(tableTest, myMock)

	// case 2, glob prok Silver
	// expected result 68
	myMock = mockForTest{
		Input:    "glob prok Silver",
		Output:   68,
		TestName: "case2",
	}
	tableTest = append(tableTest, myMock)

	// case 3, glob prok Gold
	// expected result 57800
	myMock = mockForTest{
		Input:    "glob prok Gold",
		Output:   57800,
		TestName: "case3",
	}
	tableTest = append(tableTest, myMock)

	// case 3, glob prok Iron
	// expected result 782
	myMock = mockForTest{
		Input:    "glob prok Iron",
		Output:   782,
		TestName: "case4",
	}
	tableTest = append(tableTest, myMock)

	// case 3, wood could a woodchuck chuck if a woodchuck could chuck wood
	// expected result unknown
	/*
		myMock = mockForTest{
			input:    "glob prok Iron",
			output:   -1,
			testName: "case5",
		}
		tableTest = append(tableTest, myMock)
	*/
	log.Println(tableTest)
	for idx, i := range tableTest {
		log.Println(idx, i.TestName)
		t.Run(i.TestName, func(t *testing.T) {
			actual := myUsecaseMock.Calculate(myNumberMock, i.Input)

			if actual != i.Output {
				t.Error("error output")
			} else {
				log.Println("->>", i.Output, actual)
			}
		})
	}

}
