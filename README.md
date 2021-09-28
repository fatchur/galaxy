##  Galaxy Trading 


## Problem definition
Buying and selling over the galaxy requires us to convert numbers and units.The numbers used for transactions follows similar convention to the roman numerals

Roman numerals are based on seven symbols:

| Symbol | Value | galaxy language  |
| --------------- | --------------- |--------------- |
| I | 1 | glob |
| V | 5 | prok |
| X | 10 | pish |
| L | 50 | tegj |
| C |100 | not used|



### Commodities
There are three commodities traded intergalaxy,
1. silver 
2. gold
3. iron

The prices of these commodities follow this example:<br>
```
glob glob Silver is 34
glob prok Gold is 57800
pish pish Iron is 3910
```
based on that examples, we can infer
| Commodity | Price | 
| --------------- | --------------- |
| silver | 17 | 
| gold | 57800 | 
| iron | 3910 |


### Main Task
```
We have to create a program to translate the galaxy language intu arabic number.
```
The program written in `Golang`, the foldering structure is follow
1. `app/`: main program
2. `domain/`: contains the jobs related to the task. example: trading, bring money to earth etc. But in this case we just have one activity `trading`. <br>

in trading package, we have `usecase` and `models`. <br>
`usecase` is a package for all sub-activities related to intergalaxial trading. Based on the problem, there is just one sub activity supporting `trading`, **translating the galaxy language into arabic number**. The galaxy languange follows the rule of roman number.<br>

**the translation method**
```go 
func (u Usecase) Calculate(myNumber models.RomanNumberInterface, word string) int {
	s := strings.Split(word, " ")

	totalVal := 0
	prevVal := 0
	for _, val := range s {
		totalVal, prevVal = myNumber.Logic(prevVal, totalVal, val)

		if val == "Gold" {
			totalVal = totalVal * u.GoldVal
		}
		if val == "Silver" {
			totalVal = totalVal * u.SilverVal
		}
		if val == "Iron" {
			totalVal = int(float32(totalVal) * u.IronVal)
		}
	}
	return totalVal
}

```

The `models` package in `trading` is collection of some structs or constants used for supporting the `trading` actovity.  For example:
```go
type RomanNumber struct {
	Glob int
	Prok int
	Pish int
	Tegj int
}
```


3. `models/` collection of common structs or constants supporting the main program, example error struct, jwt struct etc.

Our code is **mostly using interface as method input parameters**. The reason is for flexibility to any changes. You can introduce new galaxy number without changing the methods. 

```
A great rule of thumb for Go is accept interfaces, return structs.

–Jack Lindamood
```

### Testing scenario 
we have 2 unit test,
1. `domain/trading/usecase/usecase_test.go` <br> 
for checking the translation logic. the test case are: <br>

| Input | expected result | 
| --------------- | --------------- |
| pish tegj glob glob | 42 | 
| glob prok Silver | 68 | 
| glob prok Gold | 57800 |
| glob prok Iron | 782  |
| wood could a woodchuck chuck if a woodchuck could chuck wood | -1 (unknown) | 


2. `domain/trading/number/number_test.go` <br> 
for checking the value of two sequence numbers.
example: 
```
glob glob = 2
but 
glob pish = 9 NOT 11
```

| Input | expected result | 
| --------------- | --------------- |
| glob | 1 | 
| glob pish | 9 | 
| glob tegj | 49 |
| pish pish | 20  |