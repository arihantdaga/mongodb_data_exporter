package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Query struct {
	Db string
	Collection string
	Filter interface{}
}

var client *mongo.Client 

func openConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI("<<MongoDB Connection URI>>"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return  err
	}
	return nil
}

func checkConnection() (bool, error){
	// TODO: Check Connection.
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return false, err
	} 
	return true, nil
}

func closeConnection(ctx context.Context) {
	client.Disconnect(ctx)
}

func query(ctx context.Context, q Query){
	
	if connected, err := checkConnection(); connected != true {
	  log.Fatal(err)
	}
	// filter := bson.D{{}}

	// collection := client.Database(q.Db).Collection(q.Collection)
	// cur, err:= collection.Find(ctx, filter)

}
