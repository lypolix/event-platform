package service

import (
    "context"
    "event-platform/graph/model"
    "event-platform/internal/repository"
    "sync"
)

type EventService struct {
    Repo        *repository.EventRepository
    subscribers []chan *model.Event
    mu          sync.Mutex
}

func NewEventService(repo *repository.EventRepository) *EventService {
    return &EventService{
        Repo:        repo,
        subscribers: make([]chan *model.Event, 0),
    }
}

func (s *EventService) CreateEvent(ctx context.Context, event *model.Event) (string, error) {
    id, err := s.Repo.CreateEvent(ctx, event)
    if err == nil {
        event.ID = id
        s.mu.Lock()
        for _, sub := range s.subscribers {
            select {
            case sub <- event:
            default:
            }
        }
        s.mu.Unlock()
    }
    return id, err
}

// Получить все события
func (s *EventService) ListEvents(ctx context.Context) ([]*model.Event, error) {
    return s.Repo.GetAllEvents(ctx)
}

// Получить одно событие по ID
func (s *EventService) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
    return s.Repo.GetEventByID(ctx, id)
}

// Обновить данные события
func (s *EventService) UpdateEvent(ctx context.Context, id string, updateData map[string]interface{}) error {
    return s.Repo.UpdateEvent(ctx, id, updateData)
}

// Удалить событие
func (s *EventService) DeleteEvent(ctx context.Context, id string) error {
    return s.Repo.DeleteEvent(ctx, id)
}

// Подписаться на поток новых событий
func (s *EventService) SubscribeToEvents(ch chan *model.Event) {
    s.mu.Lock()
    s.subscribers = append(s.subscribers, ch)
    s.mu.Unlock()
}
