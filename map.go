package main

import (
	"fmt"
)

type user struct {
	name string
	number string
}

func printmap(usermap map[string]user) {

	for i,j := range usermap {
		fmt.Println(i,j)
	}
}

func insertTomap(usermap map[string]user , strarr []string , intarr []string)  {

	if len(strarr) != len(intarr) {
		panic("Wadu hekkkkkkkkkkk..")
	}

	for i:= 0 ;i<len(strarr) ; i++ {
		name := strarr[i]
		phone := intarr[i]
		usermap[name] = user {
			name: name,
			number: phone,
		}
	}
}

func main(){

	// creating map for an int array with string indexes
	// arr := make(map[string]int)
	// arr["sandeep"] = 21
	// arr["wasu"] = 19
	// arr["hekk"] = 33

	// another way
	arr := map[string]int{
		"sandeep" : 21,
		"wadu" : 19,
		"hekk" : 33,
	}

	fmt.Println(arr,len(arr))


	// map of structs

	usermap := make(map[string]user)

	strarr := []string{"wadu","suresh","rishi","sunny","adarsh","yash"}
	intarr := []string{"+919685485584","+917855784469","+918895785587","+916987854410","+918877885143","+917742174852"}


	// maps are passed by referense by default in go , so no declaration of map again 
	insertTomap(usermap,strarr,intarr)

	printmap(usermap)

	delete(usermap,"wadu")
	fmt.Println("\n")

	printmap(usermap)


}