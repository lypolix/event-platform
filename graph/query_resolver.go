package graph

import (
	"context"
	"event-platform/graph/model"
	
)


func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
    return r.UserService.ListUsers(ctx)  // Метод сервиса для получения списка пользователей
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
    return r.UserService.GetUserByID(ctx, id)
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
    return r.EventService.ListEvents(ctx)
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
    return r.EventService.GetEventByID(ctx, id)
}
