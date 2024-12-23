package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

type System struct {
	eventBus            *EventBus
	userService         *UserService
	buffer              *Buffer
	dispatcher          *ApplicationDispatcher
	statistics          *Statistics
	notificationService *NotificationService
	generator           *PoissonGenerator
	stepInterval        float64 // интервал одного шага в секундах
}

func (s *System) processNextStep() {
	// Сначала пытаемся обработать заявки из буфера
	s.dispatcher.processBuffer()

	// Симулируем работу учителей
	s.dispatcher.simulateTeachersWork()

	eventsCount := s.generator.GetEventsCountForInterval(s.stepInterval)

	// Создаем все сгенерированные заявки
	for i := 0; i < eventsCount; i++ {
		app := s.userService.CreateApplication()
		s.eventBus.Publish(Event{
			Type:      "NewApplication",
			Data:      app,
			Timestamp: time.Now(),
		})
		s.dispatcher.ProcessApplication(app)
	}
}

func (s *System) printSystemState() {
	visualizer := NewSystemVisualizer(s)
	visualizer.PrintSystemState()
}

func (s *System) printFinalStatistics() {
	s.statistics.PrintCurrentStats()
}

func (s *System) printFinalDigitsStatistics() {
	s.statistics.PrintDigitCurrentStats()
}

func NewSystem(bufferSize, teacherCount, teacherLoad int, lambda float64, stepInterval float64) *System {
	eventBus := NewEventBus()
	buffer := NewBuffer(bufferSize, eventBus)

	return &System{
		eventBus:            eventBus,
		userService:         NewUserService(eventBus),
		buffer:              buffer,
		dispatcher:          NewApplicationDispatcher(teacherCount, teacherLoad, buffer, eventBus),
		statistics:          NewStatistics(eventBus),
		notificationService: NewNotificationService(eventBus),
		generator:           NewPoissonGenerator(lambda),
		stepInterval:        stepInterval,
	}
}

func (s *System) RunStepMode() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		fmt.Println(time.Now())
		fmt.Println("\nPress Enter to continue or 'q' to quit...")

		var input string
		fmt.Scanln(&input)
		if input == "q" {
			break
		}

		screen.Clear()
		s.processNextStep()
		s.printSystemState()
	}
}

func (s *System) RunAutoMode() {
	for j := 0; j < 5; j++ {
		for i := 0; i < 100; i++ {
			s.processNextStep()
			time.Sleep(time.Second / 50)
		}
	}
	s.printFinalDigitsStatistics()
}
