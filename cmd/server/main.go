package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yashwanth-reddy909/go-crud/http"
	"github.com/yashwanth-reddy909/go-crud/http/handlers"
	"github.com/yashwanth-reddy909/go-crud/repositories/mongodb"
	"github.com/yashwanth-reddy909/go-crud/services/articles"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	articleRepoMongo := mongodb.NewArticleRepository("test", client)
    articleService := articles.NewArticleService(articleRepoMongo)
	articleHandler := handlers.NewArticleHandler(articleService)
	server := http.NewServer(articleHandler)
	server.ListenAndServe()
	
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}