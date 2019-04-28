# Exercises

## Week-01

### Hello

#### How to run

##### Option 1

`cd week1/01-hello && go run ./hello.go` //Hello, world!

##### Option 2

`cd week1/01-hello && go build ./hello.go && ./hello` //Hello, world!

## Week-02

### Trunc

#### PROMPT

Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

#### How to run

`cd week2/01-trunc && go run ./trunc.go`
`$ Please enter a Floating Point Number: 4.4` //The truncated input is: 4
`$ Please enter a Floating Point Number: 2` //The truncated input is: 2

### Findian

#### PROMPT

Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

#### Examples

The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

#### How to run

`cd week2/02-findian && go run ./findian.go`

`$ Please enter a string: ian` //Found!
`$ Please enter a string: Ian` //Found!
`$ Please enter a string: iuiygaygn` //Found!
`$ Please enter a string: "I d skd a efju N"` //Found!
`$ Please enter a string: ihhhhhn` //Not Found!
`$ Please enter a string: ina` //Not Found!
`$ Please enter a string: xian` //Not Found!

## Week-03

### Slice

#### PROMPT

Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

#### How to run

`cd week3/01-slice && go run ./slice.go`
`$ Please enter an Integer: 4` //The sorted slice is: 4
`$ Please enter an Integer: 3` //The sorted slice is: 3,4
`$ Please enter an Integer: 5` //The sorted slice is: 3,4,5
`$ Please enter an Integer: 1` //The sorted slice is: 1,3,4,5
`$ Please enter an Integer: X` //Exit