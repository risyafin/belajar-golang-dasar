package main

import "fmt"

func main (){
	name:= "samsul"
	name2:="samsul"
	name3 :="arifin"

	result := name == name2
	fmt.Println(result)
	result = name == name3
	fmt.Println(result)
	result = name != name2
	fmt.Println(result)
	result = name != name3
	fmt.Println(result)
}