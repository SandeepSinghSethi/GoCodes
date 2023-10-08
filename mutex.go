package main

import (
	"fmt"
	"sync"
)

func f1(c1 chan int , i int){
	i.
}

func f2(c2 chan int , i int){
	
}

func main(){
	c1,c2 := make(chan int),make(chan int)
	i:=0;

	go f1(c1,i)
	go f2(c2,i)

}