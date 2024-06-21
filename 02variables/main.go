package main

import "fmt"

func main(){
	// fmt.Println("variables");
	var username string = "variables" // string
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	// var age int = 23 // number or integer
	// var condition bool = true
	// var fll float64 = 100.100
	
	// fmt.Println(age) // printing age
	// fmt.Println(condition) // condition
	// fmt.Print(fll) // float
}