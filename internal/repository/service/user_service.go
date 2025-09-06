package service

import (
	"context"
	"event-platform/graph/model"
	"event-platform/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	panic("unimplemented")
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Получить пользователя по ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

// Создать пользователя
func (s *UserService) CreateUser(ctx context.Context, user *model.User) (string, error) {
	// Можно добавить валидацию или бизнес-логику здесь
	return s.Repo.CreateUser(ctx, user)
}

// Обновить пользователя
func (s *UserService) UpdateUser(ctx context.Context, id string, updateData map[string]interface{}) error {
	// Преобразуем map в bson.M
	bsonUpdate := make(map[string]interface{})
	for k, v := range updateData {
		bsonUpdate[k] = v
	}
	return s.Repo.UpdateUser(ctx, id, bsonUpdate)
}

// Удалить пользователя
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.DeleteUser(ctx, id)
}
