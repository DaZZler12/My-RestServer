package store

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateCollection creates a new collection in the specified database
func StoreCreateCollection(db *mongo.Database, collectionName string) (*mongo.Collection, error) {
	collection := db.Collection(collectionName)
	if collection == nil {
		return nil, fmt.Errorf("failed to create collection %s", collectionName)
	}
	return collection, nil
}
