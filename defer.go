package main

import(
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
	"math/rand"
	"time"
	"runtime"
)

type CInfo struct{
	Name string
	Model string
	carAge int
	color string
}

// we can assign functions to structs to do some tasks 
func (c CInfo) info() {
	fmt.Printf("CarName: %s\tCarModel: %s\tCarAge: %d\tCarColor: %s\n",c.Name,c.Model,c.carAge,c.color)
}


type CarOwner struct{
	LicenseNumber int
	CarNumber string
	Owner string
	age int
	country string
	CInfo 
	// can be used like this but i will not recommend as assignment will confuse u
	// carinfo struct{
	// 	Name string
	// 	Model string
	// 	carAge int
	// 	color string	
	// }
}


// struct name must start with caps
type Student struct{
	Firstname string
	Lastname string
	age uint8
	address string
	city string
	state string
	country string
	pincode uint32
}

func getCords() (int,int){
	return 3,4
}




func switchcase(){
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	fmt.Println("Choosing a random number [1-5]: ")

	switch num:= r.Intn(3); num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Random function pooped .. ")
	}
}

func main(){
	switchcase()
	num := 1
	defer fmt.Println("Value of num for the defer funciton : ", num)
	for i:=0;i<10;i++ {
		num *= 2
	}
	fmt.Println("Value of num within main func : ",num)


	// using class
	// can assign values based on their indices of the struct
	stud1 := Student{"Sandeep","Singh",21,"New Avenue","Lucknow","UP","India",223321}

	// creating with randomly placed values , else will be false,0,"" based on their type but assing those values with their respective variable name
	stud2 := Student{Firstname: "Sandeep",state: "MP",country: "India",age: 32}
	fmt.Println(stud1,stud2)

	// using &stud2 (copying) or &Student instantiating as a exact pointer to the structure will work but will be identified by a "&" on the terminal , so &Student == *p || or , *&Student == p
	p := &stud2
	q := &Student{Firstname: "Hi",Lastname: "Hello"}
	fmt.Println(*p,*q,"\n\n")


	// arrays : aren't resizeable 
	// a :=[10]int{} == var a [10]int
	// a := [5]int{1,2}
	// a[0] = 1
	// a[1] = 2
	// fmt.Println(a)


	// strings
	// var b [3]string
	// b := [3]string{"helloworld","wadu","here"}
	// b[0] = "hello"
	// b[1] = "world"
	// fmt.Println(b)

	// const Myage = 50
	// const Myfloat = 2.645
	// bio := "Hi i am go lang  "
	// msg := fmt.Sprintf("%s and my age is %d and i am studying for past %.1f years",
	// 	bio,
	// 	Myage,
	// 	Myfloat,
	// )
	// fmt.Println(msg)

	// getting runtime
	verInfo := fmt.Sprintf("Runtime.Version() : %s (%t)",runtime.Version(),runtime.Version(),)
	version := runtime.Version()
	version = strings.Replace(version,"go1.","",-1)
	VerbNum,_ := strconv.ParseFloat(version,64)
	fmt.Println(verInfo)

	if VerbNum >= 19 {
	} else {
		logger := log.New(os.Stderr,"[!]",0)
		logger.Fatal("Version is too old ")
	}


	// x := 100
	// fmt.Println(x)
	// x = 101
	// y := *&x
	// y += 1
	// fmt.Println(x)
	// fmt.Println(y)

	// _ is ignored
	x,_ := getCords()
	fmt.Println(x)

	// inherited structs example
	// for using named inheritance of struct we have to assign with each and every variable with its value else error will occur
	bhikarikicar := CarOwner{
		LicenseNumber: 6836969,CarNumber: "MP32NN1234",Owner: "Mr. Bhikari Kumar",age: 50,country: "IN",
 		CInfo: CInfo{
 			Name: "Maruti Suzuki", Model: "800 D",carAge: 15,color: "RED",},}
	fmt.Println(bhikarikicar)


	myinfo := CInfo{"Wagonr","1000 D",8,"WHITE"}
	myinfo.info()


}