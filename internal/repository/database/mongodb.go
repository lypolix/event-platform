package database


import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database 

func CreateCollections() error {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    usersCollection := Database.Collection("users")
    emailIndex := mongo.IndexModel{
        Keys:    bson.D{{Key: "email", Value: 1}},
        Options: options.Index().SetUnique(true),
    }

    createdAtIndex := mongo.IndexModel{
        Keys: bson.D{{Key: "created_at", Value: -1}},
    }

    if _, err := usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{emailIndex, createdAtIndex}); err != nil {
        return fmt.Errorf("failed to create users indexes: %v", err)
    }

    eventsCollection := Database.Collection("events")

    titleIndex := mongo.IndexModel{
        Keys: bson.D{{Key: "title", Value: "text"}}, 
    }

    dateTimeIndex := mongo.IndexModel{
        Keys: bson.D{{Key: "date_time", Value: 1}}, 
    }

    if _, err := eventsCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{titleIndex, dateTimeIndex}); err != nil {
        return fmt.Errorf("failed to create events indexes: %v", err)
    }

    log.Println("Collections 'users' and 'events' and indexes created successfully")
    return nil
}
