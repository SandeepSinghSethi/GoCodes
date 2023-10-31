package main

import (
	"fmt"
	"sync"
	"time"
)

type info struct{
	name string
	age int
	sync.Mutex
}

func f1(c info){
	c.Lock()
	defer c.Unlock()
	c.age++
	fmt.Println(c.age)
}

func main(){
	mapint := map[string]int{
		"a" : 1,
		"b" : 2,
		"c" : 3,
	}
	mapstring := map[string]string{
		"1" : "a",
		"2" : "b",
		"3" : "c",
	}

	fmt.Println(mapint,mapstring)

	i := info{name:"A",age:10}

	go f1(i)
	go f1(i)
	time.Sleep(time.Second+5)
	fmt.Println(i.age);

}