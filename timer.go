package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

//convert str from Input/Argument to time.Duration for timer
func argToTime(str string) time.Duration {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return time.Duration(i) 
}

func main () {
	var minute time.Duration
	var x string
	
	//check for Args
	if len(os.Args) < 2 {
		fmt.Print("Enter a number of minutes:\n>")
		fmt.Scanln(&x)
		minute = argToTime(x)
	} else {
		minute = argToTime(os.Args[1])
	}

	fmt.Println("Number of minutes:", minute)
	//timer1 := time.NewTimer(minute * time.Minute)
	timer1 := time.NewTimer(minute * time.Second)
	fmt.Println("Timer starts now")
	<-timer1.C
	fmt.Println("Timer ends")
}

