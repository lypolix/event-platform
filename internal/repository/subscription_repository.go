package repository

import (
    "context"
    "event-platform/graph/model"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionRepository struct {
    Collection *mongo.Collection
}

func NewSubscriptionRepository(collection *mongo.Collection) *SubscriptionRepository {
    return &SubscriptionRepository{Collection: collection}
}

// Пример CRUD методов ниже, можно реализовать позже по необходимости:
func (r *SubscriptionRepository) CreateSubscription(ctx context.Context, subscription *model.Subscription) (string, error) {
    res, err := r.Collection.InsertOne(ctx, subscription)
    if err != nil {
        return "", err
    }
    oid := res.InsertedID.(primitive.ObjectID)
    return oid.Hex(), nil
}
// Остальные методы: GetSubscriptionByID, ListSubscriptions и т.д. — по аналогии.
