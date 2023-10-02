package main

import (
	"fmt"
	"time"
	"math/rand"
)

func gofunc(num int,ch chan int) {
	ch <- num
}

func main(){
	ch := make(chan int)

	seed := rand.NewSource(time.Now().UnixNano())
	r:= rand.New(seed)

	go func(){
		for i:=0;i<11;i++ {
			gofunc(r.Intn(100),ch)
		}
		close(ch)
	}()
	
	for {
		val,ok := <-ch
		if !ok {
			break
		}

		fmt.Println(val)
	}

	// buffered channel
	buf := make(chan int,100)

	go func(){
		for i:=0;i<100;i++ {
			buf <- i
		}
		close(buf)
	}()


	for item := range buf {
		fmt.Println(item)
	}


	// fmt.Println(buf)


}