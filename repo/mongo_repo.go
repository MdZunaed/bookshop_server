package repo

import (
	"context"

	"github.com/MdZunaed/bookshop/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepoInterface interface {
	Create(data interface{}, ctx mongo.SessionContext) (interface{}, error)
	FindOne(id string, ctx mongo.SessionContext) (interface{}, error)
	Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error)
	Delete(id string, ctx mongo.SessionContext) (interface{}, error)
	FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error)
	Aggregate(pipeline mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]any, error)
	FindOneByKey(key string, value any, ctx mongo.SessionContext) (any, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func GetMongoRepository(dbName string, collectionName string) MongoRepoInterface {
	collection := config.GetDatabaseCollection(&dbName, collectionName)
	return &MongoRepository{
		collection: collection,
	}
}

func setupSessionContext(sessionContext mongo.SessionContext) mongo.SessionContext {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	ctx := context.Background()

	if sessionContext == nil {
		return mongo.NewSessionContext(ctx, mongo.SessionFromContext(ctx))
	}
	return sessionContext
}

func (mr *MongoRepository) Create(data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := setupSessionContext(ctx)
	result, err := mr.collection.InsertOne(sessionContext, data)
	return result, err
}

func (mr *MongoRepository) FindOne(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := setupSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := mr.collection.FindOne(sessionContext, bson.M{"_id": objId})
	var document map[string]interface{}
	if err := result.Decode(&document); err != nil {
		return nil, err
	}
	return document, nil
}

func (mr *MongoRepository) FindOneByKey(key string, value any, ctx mongo.SessionContext) (any, error) {
	sessionContext := setupSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(key)
	var result *mongo.SingleResult
	if err != nil {
		result = mr.collection.FindOne(sessionContext, bson.M{key: value})
	} else {
		result = mr.collection.FindOne(sessionContext, bson.M{key: objId})
	}
	var document map[string]any
	if err := result.Decode(&document); err != nil {
		return nil, err
	}
	return document, nil
}

func (mr *MongoRepository) Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContet := setupSessionContext(ctx)
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := mr.collection.UpdateOne(sessionContet, bson.M{"_id": objectId}, data)
	return result, err
}

func (mr *MongoRepository) Delete(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContet := setupSessionContext(ctx)
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := mr.collection.DeleteOne(sessionContet, bson.M{"_id": objectId})
	return result, err
}

func (mr *MongoRepository) FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := setupSessionContext(ctx)
	if filter == nil {
		filter = bson.M{}
	}
	cursor, err := mr.collection.Find(sessionContext, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(sessionContext)
	var results []map[string]interface{}
	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		results = append(results, document)
	}
	return results, cursor.Err()
}

func (mr *MongoRepository) Aggregate(pipeline mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]any, error) {
	sessionContext := setupSessionContext(ctx)
	cursor, err := mr.collection.Aggregate(sessionContext, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(sessionContext)
	var results []map[string]any
	for cursor.Next(sessionContext) {
		var document map[string]any
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		results = append(results, document)
	}
	return results, cursor.Err()
}
