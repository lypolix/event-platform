package service

import (
    "context"
    "event-platform/graph/model"
    "event-platform/internal/repository"
    "sync"
)

type UserService struct {
    Repo        *repository.UserRepository
    subscribers []chan *model.User
    mu          sync.Mutex
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
    return s.Repo.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
    return s.Repo.GetUserByID(ctx, id)
}

// Подписка на создание пользователя
func (s *UserService) SubscribeToUsers(ch chan *model.User) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.subscribers = append(s.subscribers, ch)
}

// Метод для оповещения всех подписчиков о новом пользователе
func (s *UserService) NotifyUserCreated(user *model.User) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for _, sub := range s.subscribers {
        sub <- user
    }
}

// Создать пользователя (+ уведомление подписчиков)
func (s *UserService) CreateUser(ctx context.Context, user *model.User) (string, error) {
    id, err := s.Repo.CreateUser(ctx, user)
    if err != nil {
        return "", err
    }
    s.NotifyUserCreated(user) // После успешного создания уведомляем подписчиков
    return id, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, updateData map[string]interface{}) error {
    return s.Repo.UpdateUser(ctx, id, updateData)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
    return s.Repo.DeleteUser(ctx, id)
}
