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
	bufferSize := flag.Int("buffer", 10, "Buffer size")
	teacherCount := flag.Int("teachers", 3, "Number of teachers")
	flag.Parse()

	// Инициализация системы
	system := NewSystem(*bufferSize, *teacherCount)

	if *isStepMode {
		system.RunStepMode()
	} else {
		system.RunAutoMode()
	}
}
