package main

import "time"
// import "fmt"
import "math/rand"

const (
	StatusNew         = "New"
	StatusInProgress  = "InProgress"
	StatusReviewing   = "Reviewing"
	StatusPassed      = "Passed"
	StatusFailed      = "Failed"
	StatusImprovement = "NeedsImprovement"
	StatusEnrolled    = "Enrolled"
)

type Application struct {
	ID        int
	UserID    int
	TestTask  *TestTask
	Status    string
	CreatedAt time.Time
}

type TestTask struct {
	ID      int
	Content string
	Status  string
	Score   int
}

func (t *TestTask) Review() bool {
	// Симулируем проверку задания с 70% шансом успеха
	return rand.Float32() < 0.7
}

type User struct {
	ID          int
	SkillLevel  string
	WaitingTime int
}

type Teacher struct {
	ID           int
	CurrentLoad  int
	MaxLoad      int
	Applications []*Application
}

func NewTeacher(id int, maxLoad int) *Teacher {
	return &Teacher{
		ID:           id,
		MaxLoad:      maxLoad,
		Applications: make([]*Application, 0),
	}
}

func (t *Teacher) completeRandomTask() *Application {
	if len(t.Applications) == 0 {
		return nil
	}

	randomIndex := rand.Intn(len(t.Applications))
	app := t.Applications[randomIndex]

	// Обрабатываем заявку
	t.processApplication(app)

	// Удаляем заявку из списка учителя
	t.Applications = append(t.Applications[:randomIndex], t.Applications[randomIndex+1:]...)
	t.CurrentLoad--

	return app
}

type NotificationService struct {
	eventBus *EventBus
}

func (t *Teacher) processApplication(app *Application) {
	app.Status = StatusReviewing

	// Проверяем тестовое задание
	if app.TestTask.Review() {
		app.Status = StatusPassed
		app.TestTask.Status = "Passed"

		// Подтверждаем зачисление
		if rand.Float32() < 0.8 { // 90% шанс подтверждения
			app.Status = StatusEnrolled
		}
	} else {
		app.Status = StatusFailed
		app.TestTask.Status = "Failed"

		if rand.Float32() < 0.4 { // 40% шанс на доработку
			app.Status = StatusImprovement
		}
	}
}

func NewNotificationService(eventBus *EventBus) *NotificationService {
	ns := &NotificationService{eventBus: eventBus}

	// Подписываемся на события
	eventChan := make(chan Event)
	eventBus.Subscribe("ApplicationBuffered", eventChan)
	eventBus.Subscribe("ApplicationRemoved", eventChan)
	eventBus.Subscribe("ApplicationProcessed", eventChan)

	go ns.handleEvents(eventChan)
	return ns
}

func (ns *NotificationService) handleEvents(eventChan chan Event) {
	for event := range eventChan {
		switch event.Type {
		case "ApplicationBuffered":
			ns.notifyUserAboutWaitingList(event.Data.(*Application))
		case "ApplicationRejected":
			ns.notifyUserAboutResult(event.Data.(*Application))
		case "ApplicationRemoved":
			ns.notifyUserAboutRemoval(event.Data.(*Application))
		case "ApplicationProcessed":
			ns.notifyUserAboutResult(event.Data.(*Application))
		}
	}
}

func (ns *NotificationService) notifyUserAboutWaitingList(app *Application) {
	// fmt.Printf("     Notification to User %d: Your application is in waiting list\n", app.UserID)
}

func (ns *NotificationService) notifyUserAboutRemoval(app *Application) {
	// fmt.Printf("     Notification to User %d: Your application was removed from waiting list\n", app.UserID)
}

func (ns *NotificationService) notifyUserAboutResult(app *Application) {
	// fmt.Printf("     Notification to User %d: Your application status is %s\n", app.UserID, app.Status)
}
