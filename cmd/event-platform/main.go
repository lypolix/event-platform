package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"event-platform/graph"
	"event-platform/internal/repository"
	"event-platform/internal/service"
)

func main() {
    // Контекст с таймаутом подключения к MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Подключение к MongoDB
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer func() {
        if err = client.Disconnect(ctx); err != nil {
            log.Fatalf("Failed to disconnect MongoDB client: %v", err)
        }
    }()

    // База данных и коллекции
    db := client.Database("event_platform")

    userRepo := repository.NewUserRepository(db.Collection("users"))
    eventRepo := repository.NewEventRepository(db.Collection("events"))

    userService := service.NewUserService(userRepo)
    eventService := service.NewEventService(eventRepo)

	subRepo := repository.NewSubscriptionRepository(db.Collection("subscriptions")) 
	subscriptionService := service.NewSubscriptionService(
		subRepo,
		userRepo,
		eventRepo,
	)

    resolver := &graph.Resolver{
        UserService:  userService,
        EventService: eventService,
		SubscriptionService: subscriptionService,
    }

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

    http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
    http.Handle("/query", srv)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
