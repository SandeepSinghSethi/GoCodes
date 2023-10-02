package main

import(
	"fmt"
)

func add(x,y int) int{
	return x+y
}

func sub(x,y int) int{
	return x-y
}

func aggr(a,b,c int ,arithmetic func(int,int) int) int {
	return arithmetic(arithmetic(a,b),c)
}

func multiply(x,y int) int {
	return x*y
}

func selfFunc(method func(int,int) int ) func(int) int {
	return func(x int) int {
		return method(x,x)
	}
}

func main(){
	fmt.Println(aggr(3,4,5,add))
	fmt.Println(aggr(3,4,5,sub))


	// function currying is a method to return a function from within a function 
	squarefunc := selfFunc(multiply)
	fmt.Println(squarefunc(4))

	addFunc := selfFunc(add)
	fmt.Println(addFunc(5))
} 