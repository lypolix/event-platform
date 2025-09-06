package service

import (
    "context"
    "event-platform/graph/model"
    "event-platform/internal/repository"
)

type EventService struct {
    Repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
    return &EventService{Repo: repo}
}

// Получить все события
func (s *EventService) ListEvents(ctx context.Context) ([]*model.Event, error) {
    return s.Repo.GetAllEvents(ctx)
}

// Получить событие по ID
func (s *EventService) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
    return s.Repo.GetEventByID(ctx, id)
}

// Создать событие
func (s *EventService) CreateEvent(ctx context.Context, event *model.Event) (string, error) {
    // Здесь можно добавить проверку, валидацию и прочее
    return s.Repo.CreateEvent(ctx, event)
}

// Обновить событие
func (s *EventService) UpdateEvent(ctx context.Context, id string, updateData map[string]interface{}) error {
    // Преобразуем map в bson.M
    bsonUpdate := make(map[string]interface{})
    for k, v := range updateData {
        bsonUpdate[k] = v
    }
    return s.Repo.UpdateEvent(ctx, id, bsonUpdate)
}

// Удалить событие
func (s *EventService) DeleteEvent(ctx context.Context, id string) error {
    return s.Repo.DeleteEvent(ctx, id)
}
