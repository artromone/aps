package main

import (
	"fmt"
)

type SystemVisualizer struct {
	system *System
}

func NewSystemVisualizer(system *System) *SystemVisualizer {
	return &SystemVisualizer{system: system}
}

func (v *SystemVisualizer) PrintSystemState() {
	fmt.Println("\n=== System State ===")
	v.printBuffer()
	v.printTeachers()
	v.printStatistics()
	fmt.Println("==================")
}

func (v *SystemVisualizer) printBuffer() {
	fmt.Print("\nBuffer State: ")
    fmt.Printf("%d/%d\n", len(v.system.buffer.applications), v.system.buffer.maxSize)
	if len(v.system.buffer.applications) > 0 {

		for i, app := range v.system.buffer.applications {
			fmt.Printf("[%d] App ID: %d, Status: %s\n", i+1, app.ID, app.Status)
		}
	}
}

func (v *SystemVisualizer) printTeachers() {
	fmt.Println("\nTeachers State:")
	for _, teacher := range v.system.dispatcher.teachers {
		fmt.Printf(" - Teacher %d: Load %d/%d\n",
			teacher.ID,
			teacher.CurrentLoad,
			teacher.MaxLoad)

		if len(teacher.Applications) > 0 {
			fmt.Println("   - Current applications:")
			for _, app := range teacher.Applications {
				fmt.Printf("     - App ID: %d, Status: %s\n", app.ID, app.Status)
			}
		}
	}
}

func (v *SystemVisualizer) printStatistics() {
	fmt.Println("\nCurrent Statistics:")
	v.system.statistics.PrintCurrentStats(1)
}

// Добавим таблицу событий
type EventTable struct {
	events []Event
}

func (et *EventTable) AddEvent(event Event) {
	et.events = append(et.events, event)
}

func (et *EventTable) PrintTable() {
	fmt.Println("\nEvent Timeline:")
	fmt.Println("Time | Type | Details")
	fmt.Println("-----|------|--------")

	for _, event := range et.events {
		fmt.Printf("%s | %s | %v\n",
			event.Timestamp.Format("15:04:05"),
			event.Type,
			event.Data)
	}
}
