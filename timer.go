package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

//convert str from Argument to time.Duration for timer
func argToTime(str string) time.Duration {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return time.Duration(i) 
}

func main () {
	minute := argToTime(os.Args[1])
	fmt.Println("Number minutes:", minute)

	//TODO: what if there is no arguments?

	timer1 := time.NewTimer(minute * time.Minute)
	fmt.Println("Timer starts now")
	<-timer1.C
	fmt.Println("Timer ends")
}

