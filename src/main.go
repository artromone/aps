package main

import (
	"flag"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 10 10 1
	isStepMode := flag.Bool("step", false, "Run in step-by-step mode")
	bufferSize := flag.Int("buffer", 3, "Buffer size")
	teacherCount := flag.Int("teachers", 3, "Number of teachers")
	teacherLoad := flag.Int("maxload", 3, "Teacher max load")
	lambda := flag.Float64("lambda", 1.0, "Poisson arrival rate (average arrivals per second)")
	stepInterval := flag.Float64("interval", 1.0, "Time interval for one step in seconds")
	flag.Parse()

	system := NewSystem(*bufferSize, *teacherCount, *teacherLoad, *lambda, *stepInterval)

	if *isStepMode {
		system.RunStepMode()
	} else {
		system.RunAutoMode()
	}
}
