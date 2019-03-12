package main

import (
	"fmt"
	"strconv"
)

//===================================
// Loop solution.
func FibLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

//===================================
// Using Recursion 
func FibRec(n int) int {
	if n <= 1 {
		return n
	}
	return FibRec(n-1) + FibRec(n-2)
}

//===================================
//===================================

//===================================
// Using Channels to get a FIB Solution.
func dup(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			println("pop x")
			x := <-in
			println("push a ", x)
			a <- x
			println("push b ", x)
			b <- x
			println("push out ", x)
			c <- x
		}
	}()
	return a, b, c
}

//===================================
func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup(x)
	go func() {
		println("push x 0")
		x <- 0
		println("push x 1")
		x <- 1
		println("pop a")
		<-a
		for {
			println("get a + b then push x ")
			x <- <-a + <-b
		}
	}()
	return out
}

//===================================
func main() {
	fmt.Println("=========================")
	fmt.Println("====  fib loop ==========")
	for i := 0; i <= 10; i++ {
		fmt.Print(strconv.Itoa(FibLoop(i)) + " ")
	}
	fmt.Println("=========================")
	fmt.Println("====  fib recursion======")

	for i := 0; i <= 10; i++ {
		fmt.Print(strconv.Itoa(FibRec(i)) + " ")
	}
	fmt.Println("=========================================")
	fmt.Println("====  fib with goroutines channels ======")
	fmt.Println("==  still working on it  ================")

	x := fib()
	for i := 0; i < 10; i++ {
		println("pop out")
		fmt.Printf("%v ", <-x)
	}

}
