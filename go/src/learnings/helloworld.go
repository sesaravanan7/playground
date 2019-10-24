package main 
//any program should have package main

import (
	"fmt"
	"runtime"
	)  //standard library
//fun main is the first function will execute

func main(){
	fmt.Println("Hello from",runtime.GOOS)  
}
