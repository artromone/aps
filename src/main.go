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
	flag.Parse()

	system := NewSystem(*bufferSize, *teacherCount, *teacherLoad)

	if *isStepMode {
		system.RunStepMode()
	} else {
		system.RunAutoMode()
	}
}
