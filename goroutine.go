package main

import (
	"fmt"
	"time"
	// "math/rand"
)

func f1(ch1 chan int,ch2 chan string){
	for{
		select{
			case val,ok := <-ch1:
				if !ok{
					fmt.Println("CH1 closed")
					return
				}
				fmt.Println(val)

			case val,ok := <-ch2:
				if !ok{
					fmt.Println("CH2 closed")
					return
				}
				fmt.Println(val)
			default:
				time.Sleep(time.Second + 2)
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
 		for i:=0;i<10;i++{
			num<-i	
 			if(i >= 8){
				num <- i
				close(num)
			} 		
			str <- string(48+i)
 		}
 		close(str)
 	}()

 	time.Sleep(time.Second + 3)
 	f1(num,str)
 	fmt.Println("hello")




}