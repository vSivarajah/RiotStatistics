package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
	"github.com/vsivarajah/RiotStatistics/producer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type dbService struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoDB(conf *config.Config) (producer.Sender, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err != nil {
		panic(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	mongoCollection := client.Database("riotdb").Collection("matchdetails")

	return &dbService{Client: client, Collection: mongoCollection}, nil

	// disconnects from mongo
}

//func (db *db) Init(ctx context.Context, config interface{}) error {
//	return nil
//}

// Replace the uri string with your MongoDB deployment's connection string.

func (db *dbService) Send(ctx context.Context, match *api.Match) error {

	matchDetailsCollection, err := db.Collection.InsertOne(ctx, match.MatchDTO)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully inserted row %s", matchDetailsCollection)
	return nil
}

func (db *dbService) Get(ctx context.Context, key int) *api.Match {
	return nil
}
