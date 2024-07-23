package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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
	timer1 := time.NewTimer(minute * time.Minute)
	//timer1 := time.NewTimer(minute * time.Second)
	fmt.Println("Timer starts now")
	<-timer1.C
	fmt.Println("Timer ends")

	/*
	All nonsense below is for making a funny sound
	when timer ends. Maybe it is titally unnecessary,
	but I already found package for it. 
	It is what it is
	*/
	
	f, err := os.Open("amogus.mp3")
	if err != nil {
		panic(err)
	}
	
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}

