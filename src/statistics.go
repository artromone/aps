package main

import "fmt"
import "time"

type Statistics struct {
	totalApplications     int
	processedApplications int
	rejectedApplications  int
	averageWaitingTime    float64
	averageWorkLoad       float64
	accumulatWorkload     float64
	eventBus              *EventBus
	processedChan         chan Event
	rejectedChan          chan Event
}

func NewStatistics(eventBus *EventBus) *Statistics {
	stats := &Statistics{
		eventBus:      eventBus,
		processedChan: make(chan Event),
		rejectedChan:  make(chan Event),
	}

	// Подписка на события через каналы
	eventBus.Subscribe("ApplicationProcessed", stats.processedChan)
	eventBus.Subscribe("ApplicationRejected", stats.rejectedChan)

	// Запуск горутин для обработки событий
	go stats.handleEvents()

	return stats
}

func (s *Statistics) handleEvents() {
	for {
		select {
		case event := <-s.processedChan:
			s.handleProcessed(event)
		case event := <-s.rejectedChan:
			s.handleRejected(event)
		}
	}
}

func (s *Statistics) handleProcessed(event Event) {
	s.totalApplications++
	s.processedApplications++
	if app, ok := event.Data.(*Application); ok {
		waitingTime := time.Since(app.CreatedAt).Seconds()
		s.averageWaitingTime = (s.averageWaitingTime*float64(s.processedApplications-1) + waitingTime) / float64(s.processedApplications)
	}
}

func (s *Statistics) handleRejected(event Event) {
	s.totalApplications++
	s.rejectedApplications++
}

func (s *Statistics) PrintCurrentStats(steps int) {
	fmt.Printf("\nCurrent Statistics:\n")
	fmt.Printf("Total Applications: %d\n", s.totalApplications)
	fmt.Printf("Processed: %d\n", s.processedApplications)
	fmt.Printf("Rejected: %d\n", s.rejectedApplications)
	fmt.Printf("Average Waiting Time: %.4f seconds\n", s.averageWaitingTime)
	fmt.Printf("Average WorkLoad: %.4f percent\n", s.averageWorkLoad/float64(steps)*100)
}

func (s *Statistics) PrintDigitCurrentStats(steps int) {
	fmt.Printf("%d\n", s.totalApplications)
	fmt.Printf("%d\n", s.processedApplications)
	// fmt.Printf("%d\n", s.rejectedApplications)
	fmt.Printf("%.4f\n", s.accumulatWorkload/float64(steps)*100)
	fmt.Printf("%.4f\n", s.averageWaitingTime)
}
