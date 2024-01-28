package mongodb

import (
	"fmt"

	"github.com/yashwanth-reddy909/go-crud/models"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepository struct {
	database string
	client *mongo.Client
}

func NewArticleRepository(database string, client *mongo.Client) *ArticleRepository{
	return &ArticleRepository{database: database, client: client}
}

const collectionName string = "articles"

func (a *ArticleRepository) FindAll() ([]*models.Article, error) {
	collection := a.client.Database(a.database).Collection(collectionName)
	cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(context.Background())

	var articles []*models.Article

    if err := cur.All(context.Background(), &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleRepository) Create(article *models.Article) (bool, error) {
	collection := a.client.Database(a.database).Collection(collectionName)
	fmt.Println("Inserting with id", article.Id)
	_, err := collection.InsertOne(context.Background(), article)
    if err != nil {
        return false, err
    }
	return true, nil	
}

func (a *ArticleRepository) Update(id string, article *models.ArticleModelWithoutID) (int, error) {
	collection := a.client.Database(a.database).Collection(collectionName)
	updateFilter := bson.M{"_id": id}
	update := bson.M{"$set": article} 
	updatedResult, err := collection.UpdateOne(context.Background(), updateFilter, update)
    if err != nil {
        return -1, err
    }
	return int(updatedResult.ModifiedCount), nil	
}

func (a *ArticleRepository) Delete(id string) (int, error) {
	collection := a.client.Database(a.database).Collection(collectionName)
	fmt.Println("Deleting Article Document with id", id)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return -1, err
    } else {
		return int(deleteResult.DeletedCount), nil 
	}
}

