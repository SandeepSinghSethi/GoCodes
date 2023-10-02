package main

import (
	"fmt"
)

type car struct {
	color string
}

func (c *car) setcolor(color string){
	c.color = color
}

// for demonstrating pointer's requirement in structs
func (c car) setc(color string){
	c.color = color
}

func main(){
	// x := 5
	// y := x
	// z := &x
	// *z = 10
	// a := z
	// b := *z
	// fmt.Printf("%p , %p , %p \n%x , %x\n",x,y,z,a,b)

	// mystring := "hllo"
	// var p *string = &mystring
	// *p = "world"

	// fmt.Printf("%s %s",mystring,*p)

	// nil pointers
	// var nilpointer *int
	// fmt.Println(*nilpointer) // will panic as dereferencing null pointer

	// using pointers with structs
	withptr := car{color:"white"}
	withptr.setcolor("blue")
	fmt.Println(withptr.color)


	withoutptr := car{color:"white"}
	withoutptr.setc("blue")
	fmt.Println(withoutptr.color) // this will be white as setc will not set the actual structure

}
