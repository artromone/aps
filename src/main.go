package main

import (
	"flag"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Параметры запуска
	isStepMode := flag.Bool("step", false, "Run in step-by-step mode")
	bufferSize := flag.Int("buffer", 1, "Buffer size")
	teacherCount := flag.Int("teachers", 1, "Number of teachers")
	teacherLoad := flag.Int("maxload", 5, "Teacher max load")
	flag.Parse()

	// Инициализация системы
	system := NewSystem(*bufferSize, *teacherCount, *teacherLoad)

	if *isStepMode {
		system.RunStepMode()
	} else {
		system.RunAutoMode()
	}
}
