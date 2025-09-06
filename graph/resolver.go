package graph

import (
	"event-platform/graph/model"
	"event-platform/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Events map[string]*model.Event
    Users  map[string]*model.User
	EventsCollection *mongo.Collection
	UserService  *service.UserService
    EventService *service.EventService
}
