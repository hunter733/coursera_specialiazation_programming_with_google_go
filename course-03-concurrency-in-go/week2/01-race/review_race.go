package main

/*
 Foo.Race creates a race condition by concurrently increasing and
 decreasing the value of Foo.

 Foo.Deposit is compiled to 3 machine code instructions
 A. read the value of f
 B.	Add f and amount
 C. Set the new value of f to be the result of B.

 A race condition happens when Foo.Withdraw is executed after A. is
 executed but before C is executed.

 Similarly:
 Foo.Withdraw is compiled to 3 machine code instructions
 A. read the value of f
 B.	Subtracts amount from f
 C. Set the new value of f to be the result of B.

 A race condition happens when Foo.Deposit is executed after A. is
 executed but before C is executed.
*/

type Foo float64

func (f *Foo) Deposit(amount Foo) {
	*f = *f + amount
}

func (f *Foo) Withdraw(amount Foo) {
	// race
	*f = *f - amount
}

func (f *Foo) Race() {
	go func() {
		for i := 100.; i > 0; i-- {
			// race
			f.Deposit(Foo(i))
		}
	}()

	go func() {
		for i := 1.; i < 100; i++ {
			// race
			f.Withdraw(Foo(i))
		}
	}()
}

func main() {
	var foo Foo
	foo.Race()
}
