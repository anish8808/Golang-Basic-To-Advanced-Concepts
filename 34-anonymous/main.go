package main

import "fmt"

//anonymous funcs that are called without created the name(but we need function)

//anonymous funcs that are called without created the name(but we need function)

func main() {

	//no aurgements no retunr
	func( /*input paramtere*/ ) {
		//function body

		fmt.Println("Hello World") // everthing like normal only name won't be there
	}( /*function executor*/ ) // for executing we need to call executor

	//with one aurgement with no return
	func(msg string) {
		fmt.Println(msg)
	}("Hello World")

	f := func(msg string) string {
		return msg
	}

	sum := func(a, b int) int {
		return a + b
	}

	sum2 := func(a, b int) int {
		return a + b
	}(10, 20)

	str := f("Hello World-2")
	fmt.Println(str)

	fmt.Println("sum1", sum(10, 20))
	fmt.Println("sum2", sum2)

	add1 := Calc(10, 20, sum) //http handlers , middlewares use this
	fmt.Println("add1", add1)

	//for subsctraction

	sub1 := Calc(10, 20, func(a, b int) int {
		return a - b
	})

	fmt.Println("sub1", sub1)

	multi := Calc(10, 20, func(a, b int) int {
		return a * b
	})

	fmt.Println(multi)
}

func Calc(a, b int, f func(int, int) int) int { //here we are passing func as paramtere with particular signature
	return f(a, b)
}
