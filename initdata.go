package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type State struct {
	Name string `json:"name"`
}

var client *mongo.Client
var statesCollection *mongo.Collection

func main() {
	initDB()
	defer client.Disconnect(context.Background())

	states := readStates()
	insertStates(states)
}

func initDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	statesCollection = client.Database("test").Collection("states")
}

func readStates() []State {
	data, err := ioutil.ReadFile("main/states.json")
	if err != nil {
		log.Fatalf("Error reading states.json: %v", err)
	}

	var states []State
	err = json.Unmarshal(data, &states)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return states
}

func insertStates(states []State) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var documents []interface{}
	for _, state := range states {
		documents = append(documents, state)
	}

	_, err := statesCollection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatalf("Error inserting states into MongoDB: %v", err)
	}
	fmt.Println("States have been successfully initialized in the database!")
}
