package main

import (
	"fmt"
	"time"
	// "math/rand"
)

func f1(ch1 <-chan int,ch2 <-chan string){
	for{
		select{
			case val,ok := <-ch1:
				if !ok{
					ch1 = nil
					fmt.Println("CH1 closed")
				}else{
					fmt.Println(val," <- Ch1")
				}
			case val,ok := <-ch2:
				if !ok{
					fmt.Println("CH2 closed")
					ch2 = nil
				}else{
					fmt.Println(val, " <- Ch2")
				}
		}
		if ch1 == nil && ch2 == nil{
			break
		}
	}
}


// func gofunc(num int,ch chan int) {
// 	ch <- num
// }

func main(){
	// ch := make(chan int)

	// seed := rand.NewSource(time.Now().UnixNano())
	// r:= rand.New(seed)

	// go func(){
	// 	for i:=0;i<11;i++ {
	// 		gofunc(r.Intn(100),ch)
	// 	}
	// 	close(ch)
	// }()
	
	// for {
	// 	val,ok := <-ch
	// 	if !ok {
	// 		break
	// 	}

	// 	fmt.Println(val)
	// }

	// // buffered channel
	// buf := make(chan int,100)

	// go func(){
	// 	for i:=0;i<100;i++ {
	// 		buf <- i
	// 	}
	// 	close(buf)
	// }()


	// for item := range buf {
	// 	fmt.Println(item)
	// }


 	// fmt.Println(buf)
 	// writing to a channel will panic
 	num , str := make(chan int),make(chan string)


 	go func(){
 		for i:=0;i<8;i++{	
 				num <-i
 		}
 		time.Sleep(time.Millisecond + 500)
 		close(num)
 	}()

 	go func(){
 		for i:=0;i<10;i++{
			str <- string(48+i)
 		}
 		time.Sleep(time.Millisecond + 500)
 		close(str)
 	}()

 	time.Sleep(time.Second)
 	f1(num,str)
 	fmt.Println("hello")




}	