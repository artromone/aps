package main

import "time"
import "fmt"

type UserService struct {
	eventBus *EventBus
	nextID   int
}

func NewUserService(eventBus *EventBus) *UserService {
	return &UserService{
		eventBus: eventBus,
		nextID:   1,
	}
}

func (us *UserService) CreateApplication() *Application {
	app := &Application{
		ID:        us.nextID,
		UserID:    us.nextID,
		Status:    "New",
		CreatedAt: time.Now(),
		TestTask: &TestTask{
			ID:      us.nextID,
			Content: fmt.Sprintf("Test task #%d", us.nextID),
			Status:  "New",
		},
	}
	us.nextID++
	return app
}
