package mongo_database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnection struct {
	client   *mongo.Client
	database *mongo.Database
}

type MongoConnectionConfiguration struct {
	name string
}

var clients = map[string]*MongoConnection{}

func (db *MongoConnection) Init(configurations []MongoConnectionConfiguration) {
	for i := 0; i < len(configurations); i++ {
		if clients[configurations[i].name] == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
			defer func() {
				if err = client.Disconnect(ctx); err != nil {
					panic(err)
				}
			}()
			err = client.Ping(ctx, readpref.Primary())
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
			database := client.Database("testing")
			clients[configurations[i].name].client = client
			clients[configurations[i].name].database = database
		}
	}
}

func GetConnection(name string) *MongoConnection {
	return clients[name]
}

func GetClient(name string) *mongo.Client {
	return GetConnection(name).client
}

func GetDatabase(name string) *mongo.Database {
	return GetConnection(name).database
}
