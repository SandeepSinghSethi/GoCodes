package main

import (
	"fmt"
)

// generics are like templates

func slice[T any](s []T){
	mid := len(s)/2
	fmt.Println(s[:mid],s[mid:])
}

func main(){
	a := []string{"A","B","C","D"}
	b := []int{1,2,3,4}
	slice(a)
	slice(b)
}