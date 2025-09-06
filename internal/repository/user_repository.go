package repository

import (
    "context"
    "event-platform/graph/model"
    "fmt"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
   
)

type UserRepository struct {
    Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
    return &UserRepository{Collection: collection}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, fmt.Errorf("invalid user ID: %w", err)
    }

    var user model.User
    err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (string, error) {
    res, err := r.Collection.InsertOne(ctx, user)
    if err != nil {
        return "", fmt.Errorf("failed to create user: %w", err)
    }
    insertedID := res.InsertedID.(primitive.ObjectID).Hex()
    return insertedID, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, updateData bson.M) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid user ID: %w", err)
    }

    filter := bson.M{"_id": objID}
    update := bson.M{"$set": updateData}
    res, err := r.Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }
    if res.MatchedCount == 0 {
        return fmt.Errorf("user not found")
    }
    return nil
}


func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid user ID: %w", err)
    }

    res, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    if res.DeletedCount == 0 {
        return fmt.Errorf("user not found")
    }
    return nil
}
