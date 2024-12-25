package main

import "time"
// import "fmt"
import "math/rand"

type ApplicationDispatcher struct {
	teachers []*Teacher
	eventBus *EventBus
	buffer   *Buffer
}

func (d *ApplicationDispatcher) processBuffer() {
	if len(d.buffer.applications) == 0 {
		return
	}

	teacher := d.findAvailableTeacher()
	if teacher != nil {
		// Берем первую заявку из буфера
		app := d.buffer.applications[0]
		d.buffer.applications = d.buffer.applications[1:]

		d.assignToTeacher(app, teacher)

		d.eventBus.Publish(Event{
			Type:      "ApplicationTakenFromBuffer",
			Data:      app,
			Timestamp: time.Now(),
		})
	}
}

func (d *ApplicationDispatcher) simulateTeachersWork() float64 {

	averageLoad := 0.0

	for _, teacher := range d.teachers {
		// С некоторой вероятностью учитель завершает задачу
		if teacher.CurrentLoad > 0 && rand.Float32() < 0.3 { // 30% шанс завершения задачи
			if completedApp := teacher.completeRandomTask(); completedApp != nil {
				if completedApp.TestTask.Status == "Passed" {
					d.eventBus.Publish(Event{
						Type:      "ApplicationProcessed",
						Data:      completedApp,
						Timestamp: time.Now(),
					})
				}
				if completedApp.TestTask.Status == "Failed" {
					d.eventBus.Publish(Event{
						Type:      "ApplicationRejected",
						Data:      completedApp,
						Timestamp: time.Now(),
					})
				}
			}

		}

		averageLoad += float64(teacher.CurrentLoad) / float64(teacher.MaxLoad)
	}

	return averageLoad / float64(len(d.teachers))
}

func (d *ApplicationDispatcher) findAvailableTeacher() *Teacher {
	for _, teacher := range d.teachers {
		if teacher.CurrentLoad < teacher.MaxLoad {
			return teacher
		}
	}
	return nil
}

func (d *ApplicationDispatcher) assignToTeacher(app *Application, teacher *Teacher) {
	teacher.CurrentLoad++
	teacher.Applications = append(teacher.Applications, app)
	app.Status = "InProgress"

	d.eventBus.Publish(Event{
		Type:      "ApplicationAssigned",
		Data:      app,
		Timestamp: time.Now(),
	})
}

func NewApplicationDispatcher(teacherCount, teacherLoad int, buffer *Buffer, eventBus *EventBus) *ApplicationDispatcher {
	teachers := make([]*Teacher, teacherCount)
	for i := 0; i < teacherCount; i++ {
		teachers[i] = NewTeacher(i+1, teacherLoad)
	}

	return &ApplicationDispatcher{
		teachers: teachers,
		eventBus: eventBus,
		buffer:   buffer,
	}
}

func (d *ApplicationDispatcher) ProcessApplication(app *Application) {
	teacher := d.findAvailableTeacher()
	if teacher == nil {
		d.buffer.Add(app)
		return
	}

	d.assignToTeacher(app, teacher)
}
