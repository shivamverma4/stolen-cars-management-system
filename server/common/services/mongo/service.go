package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	// "gotham/internal/logs"
)

func getCollection(collectionName string, DBName string) (collection *mongo.Collection) {
	var err error
	client, err = mongo.Connect(ctx, fmt.Sprintf("mongodb://%s:%d", dbHost, dbPort))
	if err != nil {
		// logs.Error("mongo new client failure ", err)
	}

	collection = client.Database(DBName).Collection(collectionName)
	return
}

func InsertOne(collectionName string, document interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := getCollection(collectionName, dbName)
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		// logs.Error("mongo insert failure ", collection, err)
		return err
	}
	return nil
}

func UpdateOne(collectionName string, filter, document interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := getCollection(collectionName, dbName)
	upsert := true
	updateOptions := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := collection.UpdateOne(ctx, filter, document, &updateOptions)
	if err != nil {
		return err
		// logs.Error("mongo update failure ", collection, err)
	}
	return nil
}

func FindOne(collectionName string, filter bson.D) (*mongo.SingleResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := getCollection(collectionName, dbName)
	result := collection.FindOne(ctx, filter)
	return result, nil
}

func FindAll(collectionName string, filter bson.D) ([]interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := getCollection(collectionName, dbName)
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var list = make([]interface{}, 0)

	defer result.Close(ctx)
	for result.Next(ctx) {
		var res interface{}
		err := result.Decode(&res)
		if err != nil {
			fmt.Println("error : ", err)
			return nil, err
		}

		list = append(list, res)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func GetAggregate(collectionName string, matchStage bson.D, groupStage bson.D, DBName string) *mongo.Cursor {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := getCollection(collectionName, DBName)
	defer client.Disconnect(ctx)
	showLoadedStructCursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		fmt.Println("Error : ", err)
		// logs.Error("mongo get aggregate failure ", collection, err)
		return nil
	}
	showLoadedStructCursor.Next(ctx)
	return showLoadedStructCursor
}
