# Exercises

## Week-01

### Moore's Law

#### PROMPT

Define Moore’s law and explain why it has now stopped being true. Be sure to describe all of the physical limitations that have prevented Moore’s law from continuing to be true.

## Week-02

### Race Condition

#### PROMPT

Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

#### How to run

##### Option 1

`cd week2/01-race && go run ./race.go`
`cd week2/01-race && go run -race ./race.go`

##### Option 2

`cd week2/01-race && go build ./race.go && ./race`

## Week-03

### Concurrent Sort

#### PROMPT

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

#### How to run

##### Option 1

`cd week3/01-concurrentsort && go run ./sort.go`

##### Option 2

`cd week3/01-concurrentsort && go build ./sort.go && ./sort`

## Week-04

### Dining Philosophers Problem

#### PROMPT

Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

#### How to run

##### Option 1

`cd week4/01-diningphilosophers && go run ./diningphilosophers.go`

##### Option 2

`cd week4/01-diningphilosophers && go build ./diningphilosophers.go && ./diningphilosophers`