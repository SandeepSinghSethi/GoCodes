package main

import (
	"fmt"
	"math/rand"
	"time"
)

// types :
// int int8 int16 int32 int64
// uint unit8 uint16 uint32 uint64
// bool 
// string
// byte (same as uint8)
// rune (same as int32 for unicode char)
// float32 float64
// complex64 complex128

// const cannot use := , have to use = , and have to start with a Capital letter
const Myvar = "hi"



func init(){
	fmt.Println("Running the go file , I am init function ")
}

func mul(x int , y int) int {
	return x*y
}

func addtwo(x ,y int) (int,int){
	return x+2,y+2
}

func namedreturn(x int,y int) (retx int,rety int){
	retx = x+2
	rety = y+3
	return retx,rety
}

// var can be used outside and inside of a function but (:=) can only be used inside a func
var (
	date,year int = 10,2023
	month string = "March"
	leap bool = true
)

func casting() {
	var a,b int = 20,30
	var div float32 = float32(a)/float32(b)
	// cast
	var divint int = int(div)

	fmt.Println(div,divint)
}

func rawstring(){
	// var str string = "helloworld" 
	var str string =
	`helloworld ""
	echo hello
	"WAdU"
	`
	fmt.Print(str)
}

func testswitch(){
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	fmt.Println("Choosing a random number : ")
	num := r.Intn(30)
	fmt.Println(num)
}

func ten_random_numbers(){
	seed := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(seed)

	for i:=0;i<10;i++ {
		fmt.Print(r.Intn(100)," ,")		
	}
}


func main(){
	// fmt.Println("Helloworld")
	// fmt.Println(mul(2,6))
	// fmt.Println(addtwo(4,5))


	// // with int declaration it is clear that a and b are int
	// var a,b int = 1,2
	// // without any declaration of type it can be anything
	// var c,d,e = "hello" , 1 , true
	// var f,g int // auto set to 0
	// var h,i bool // auto set to false
	// var j,k string // auto set to ""

	// // can use := to declare a variable without var
	// date,month,year,leap := 12,"July",2022,false
	// fmt.Println(a,b,c,d,e,f,g,h,i,j,k)
	// fmt.Println(date,month,year,leap)

	// casting()

	// const One = 1
	// fmt.Println(Myvar,One)

	// rawstring()

	for i := 0 ; i < 20 ; i++ {
		fmt.Print(i," ")
	}

	// acts as a while loop
	// for{
	// 	fmt.Println("hellworld")
	// }

	// var a,b int = 0,0
	// a += 1
	// fmt.Println(a,b)


	// lexical scoping issues
	var i,sum = 0,0
	for ;i < 20 ; i++ {
		sum += i
	}

	// this will not work , same for the if 
	// for j := 0;j<20;j++{
	// 	sum += j
	// }
	// fmt.Println(j) // error

	fmt.Println(sum,i)

	var a = 100
	if a < 100{
		fmt.Println("Not 100")
	} else if a == 100 {
		fmt.Println("Is 100")
	} else {
		fmt.Println("Wadu")
	}

	// testswitch()
	ten_random_numbers()
}
