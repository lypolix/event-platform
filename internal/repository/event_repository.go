package repository

import (
    "context"
    "event-platform/graph/model"
    "fmt"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
    Collection *mongo.Collection
}

func NewEventRepository(collection *mongo.Collection) *EventRepository {
    return &EventRepository{Collection: collection}
}

func (r *EventRepository) GetAllEvents(ctx context.Context) ([]*model.Event, error) {
    cursor, err := r.Collection.Find(ctx, bson.D{})
    if err != nil {
        return nil, fmt.Errorf("failed to find events: %w", err)
    }
    defer cursor.Close(ctx)

    var events []*model.Event
    for cursor.Next(ctx) {
        var event model.Event
        if err := cursor.Decode(&event); err != nil {
            return nil, fmt.Errorf("failed to decode event: %w", err)
        }
        events = append(events, &event)
    }

    if err = cursor.Err(); err != nil {
        return nil, fmt.Errorf("cursor error: %w", err)
    }

    return events, nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, fmt.Errorf("invalid event ID: %w", err)
    }

    var event model.Event
    err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
    if err != nil {
        return nil, fmt.Errorf("event not found: %w", err)
    }

    return &event, nil
}

// Создать новое событие
func (r *EventRepository) CreateEvent(ctx context.Context, event *model.Event) (string, error) {
    res, err := r.Collection.InsertOne(ctx, event)
    if err != nil {
        return "", fmt.Errorf("failed to create event: %w", err)
    }

    insertedID := res.InsertedID.(primitive.ObjectID).Hex()
    return insertedID, nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, id string, updateData bson.M) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid event ID: %w", err)
    }

    filter := bson.M{"_id": objID}
    update := bson.M{"$set": updateData}
    res, err := r.Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return fmt.Errorf("failed to update event: %w", err)
    }

    if res.MatchedCount == 0 {
        return fmt.Errorf("event not found")
    }

    return nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid event ID: %w", err)
    }

    res, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        return fmt.Errorf("failed to delete event: %w", err)
    }

    if res.DeletedCount == 0 {
        return fmt.Errorf("event not found")
    }

    return nil
}