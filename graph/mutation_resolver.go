package graph

import (
    "context"
    "event-platform/graph/model"
    "time"
)

// Пример мутационного резолвера для создания пользователя и события

func (r *mutationResolver) CreateUser(ctx context.Context, input model.User) (*model.User, error) {
    // Создаём структуру User, заполняя поля из input и системные поля
    user := &model.User{
        // Поля из input
        Name:     input.Name,
        Email:    input.Email,
        //Password: input.Password,  // если есть, например, пароль

        // Системные поля
        CreatedAt: time.Now().UTC(),
    }

    // Вызываем сервис для сохранения пользователя и получения ID
    id, err := r.UserService.CreateUser(ctx, user)
    if err != nil {
        return nil, err
    }
    user.ID = id

    return user, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, title string, description *string, dateTime string) (*model.Event, error) {
    event := &model.Event{
        Title:       title,
        Description: description,
        DateTime:    dateTime,
        CreatedAt:   time.Now().UTC(),
    }

    id, err := r.EventService.CreateEvent(ctx, event)
    if err != nil {
        return nil, err
    }
    event.ID = id

    return event, nil
}
