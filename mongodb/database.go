package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConnection struct {
	client        *mongo.Client
	database      *mongo.Database
	ctx           context.Context
	configuration MongoDBConfiguration
}

var mongoDBConnections = map[string]*MongoDBConnection{}

func connectMongoDBConnection(configuration MongoDBConfiguration) bool {
	var databaseName = configuration.Name

	// if we have not seen/intialozed the connection
	// before then enter here and do so
	if mongoDBConnections[databaseName] == nil {
		// build context timeout for connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		// connect to server
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configuration.URI).SetAuth(options.Credential{
			Username: configuration.Username,
			Password: configuration.Password,
		}))
		if err != nil {
			panic(err)
		}

		// defer disconnetion that may happen
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()

		// make sure we are connected by pinging the server
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
			panic(err)
		}

		// connect to database and any other essentials
		database := client.Database(databaseName)

		// store necessary data for connection
		mongoDBConnections[databaseName] = &MongoDBConnection{
			client,
			database,
			ctx,
			configuration,
		}
	}

	return true
}

func Connect(configurations []MongoDBConfiguration) {
	for i := 0; i < len(configurations); i++ {
		connectMongoDBConnection((configurations[i]))
	}
}

func getMongoDBConnection(name string) *MongoDBConnection {
	var connection = mongoDBConnections[name]
	// make sure we are connected by pinging the server
	var err = connection.client.Ping(connection.ctx, readpref.Primary())
	if err = connection.client.Disconnect(connection.ctx); err != nil {
		connectMongoDBConnection(connection.configuration)
	}
	return mongoDBConnections[name]
}

func GetClient(name string) *mongo.Client {
	return getMongoDBConnection(name).client
}

func GetDatabase(name string) *mongo.Database {
	return getMongoDBConnection(name).database
}

func GetCtx(name string) context.Context {
	return getMongoDBConnection(name).ctx
}

func Shutdown(name string) {
	if err := GetClient(name).Disconnect(GetCtx(name)); err != nil {
		log.Fatal(err)
	}
}

func ShutdownAll() {
	for k := range mongoDBConnections {
		Shutdown(k)
	}
}
