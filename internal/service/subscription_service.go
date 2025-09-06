package service

import (
    "context"
    "event-platform/graph/model"
    "event-platform/internal/repository"
    "fmt"
)

// SubscriptionService реализует методы работы с подписками.
type SubscriptionService struct {
    Repo      *repository.SubscriptionRepository
    UserRepo  *repository.UserRepository
    EventRepo *repository.EventRepository
}

// Конструктор для внедрения всех зависимостей!
func NewSubscriptionService(
    repo *repository.SubscriptionRepository,
    userRepo *repository.UserRepository,
    eventRepo *repository.EventRepository,
) *SubscriptionService {
    return &SubscriptionService{
        Repo:      repo,
        UserRepo:  userRepo,
        EventRepo: eventRepo,
    }
}

// Создать подписку (простой CRUD)
func (s *SubscriptionService) CreateSubscription(ctx context.Context, sub *model.Subscription) (string, error) {
    return s.Repo.CreateSubscription(ctx, sub)
}

// Подписка пользователя на событие
func (s *SubscriptionService) SubscribeToEvent(ctx context.Context, eventID string, userID string) (*model.Subscription, error) {
    // Получаем пользователя по userID
    user, err := s.UserRepo.GetUserByID(ctx, userID)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }

    // Получаем событие по eventID
    event, err := s.EventRepo.GetEventByID(ctx, eventID)
    if err != nil {
        return nil, fmt.Errorf("event not found: %w", err)
    }

    // Создаём подписку
    sub := &model.Subscription{
        Subscriber: user,
        Event:      event,
        // SubscribedToUser: nil, если нужно добавить - реализуй по аналогии
    }

    // Сохраняем в БД
    id, err := s.Repo.CreateSubscription(ctx, sub)
    if err != nil {
        return nil, fmt.Errorf("failed to create subscription: %w", err)
    }
    sub.ID = id
    return sub, nil
}

func (s *SubscriptionService) SubscribeToUser(ctx context.Context, userID, subscriberID string) (*model.Subscription, error) {
    // Получаем пользователя, на которого подписываемся
    targetUser, err := s.UserRepo.GetUserByID(ctx, userID)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    // Получаем подписчика
    subscriber, err := s.UserRepo.GetUserByID(ctx, subscriberID)
    if err != nil {
        return nil, fmt.Errorf("subscriber not found: %w", err)
    }
    // Создаём подписку
    sub := &model.Subscription{
        Subscriber:       subscriber,
        SubscribedToUser: targetUser,
    }
    // Сохраняем её в БД
    id, err := s.Repo.CreateSubscription(ctx, sub)
    if err != nil {
        return nil, fmt.Errorf("failed to create subscription: %w", err)
    }
    sub.ID = id
    return sub, nil
}

