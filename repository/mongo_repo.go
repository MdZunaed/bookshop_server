package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func setupSessionContext(sessionContext mongo.SessionContext) mongo.SessionContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if sessionContext == nil {
		return mongo.NewSessionContext(ctx, mongo.SessionFromContext(ctx))
	}
	return sessionContext
}

func (mr *MongoRepository) Read(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := setupSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := mr.collection.FindOne(sessionContext, bson.M{"_id": objId})
	var document map[string]interface{}
	if err := result.Decode(document); err != nil {
		return nil, err
	}
	return document, nil
}
