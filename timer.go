package main

import (
	"fmt"
	//"time"
	"os"
)

func main () {
	args := os.Args[1:]
	firstArg := os.Args[1]
	fmt.Println("All arguments:", args)
	fmt.Println("First Argument:", firstArg)
	
	//TODO: what if there is no arguments?

	/*
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Timer starts now")
	<-timer1.C
	fmt.Println("Timer ends")
	*/
}

