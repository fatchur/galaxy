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
