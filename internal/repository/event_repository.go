
package repository

import (
    "context"
    "event-platform/graph/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
    Collection *mongo.Collection
}

func NewEventRepository(col *mongo.Collection) *EventRepository {
    return &EventRepository{Collection: col}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *model.Event) (string, error) {
    res, err := r.Collection.InsertOne(ctx, event)
    if err != nil {
        return "", err
    }
    oid := res.InsertedID.(primitive.ObjectID)
    return oid.Hex(), nil
}

func (r *EventRepository) GetAllEvents(ctx context.Context) ([]*model.Event, error) {
    cursor, err := r.Collection.Find(ctx, bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    var result []*model.Event
    for cursor.Next(ctx) {
        var evt model.Event
        if err := cursor.Decode(&evt); err != nil {
            return nil, err
        }
        result = append(result, &evt)
    }
    return result, nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
    oid, _ := primitive.ObjectIDFromHex(id)
    var evt model.Event
    if err := r.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&evt); err != nil {
        return nil, err
    }
    return &evt, nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, id string, updateData bson.M) error {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    filter := bson.M{"_id": oid}
    update := bson.M{"$set": updateData}
    res, err := r.Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return mongo.ErrNoDocuments
    }
    return nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, id string) error {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    res, err := r.Collection.DeleteOne(ctx, bson.M{"_id": oid})
    if err != nil {
        return err
    }
    if res.DeletedCount == 0 {
        return mongo.ErrNoDocuments
    }
    return nil
}

