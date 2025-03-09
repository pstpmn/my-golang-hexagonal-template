package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IMongoDb interface {
	Connect(uri string) (*mongo.Client, error)
	Ping(client mongo.Client) error
	CreateIndex(client *mongo.Client, dbname string)
}

type m struct{}

func (m *m) CreateIndex(client *mongo.Client, dbname string) {
	// users collection
	m.createIndex(client, dbname, "users", "email", false)
}

func NewMongo() IMongoDb {
	return &m{}
}

func (m m) Connect(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return client, err
	}

	return client, nil
}

// Ping implements IMongoDb.
func (*m) Ping(client mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	return err
}

// CreateIndex creates an index on the specified field.
func (*m) createIndex(client *mongo.Client, dbName, collectionName, fieldName string, isUnique bool) error {
	collection := client.Database(dbName).Collection(collectionName)

	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: fieldName, Value: 1}},
		// 1 for ascending order, -1 for descending order
		Options: options.Index().SetUnique(isUnique),
	}

	// Access the index manager and create the index
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
