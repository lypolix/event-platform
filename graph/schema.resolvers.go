package graph

import (
    "context"
    "event-platform/graph/model"
   
)

// Реализация мутаций

func (r *mutationResolver) SubscribeToEvent(ctx context.Context, eventID string) (*model.Subscription, error) {
    // Реализация подписки на событие - пример-заглушка
    return nil, nil
}

func (r *mutationResolver) SubscribeToUser(ctx context.Context, userID string) (*model.Subscription, error) {
    // Реализация подписки на пользователя - пример-заглушка
    return nil, nil
}

// Реализация подписок (пример-заглушки)
func (r *subscriptionResolver) ID(ctx context.Context) (<-chan string, error) {
    return nil, nil
}

func (r *subscriptionResolver) Subscriber(ctx context.Context) (<-chan *model.User, error) {
    return nil, nil
}

func (r *subscriptionResolver) Event(ctx context.Context) (<-chan *model.Event, error) {
    return nil, nil
}

func (r *subscriptionResolver) SubscribedToUser(ctx context.Context) (<-chan *model.User, error) {
    return nil, nil
}

// Методы возвращающие реализации резолверов интерфейсов
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
