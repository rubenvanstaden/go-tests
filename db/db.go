package db

import (
	"log"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Status      string             `json:"status" bson:"status"`
	Data        string             `json:"data" bson:"data"`
	Description string             `json:"description" bson:"description"`
}

type DB struct {
    *mongo.Collection   
}

func Open(url string) (*DB, error) {

	var ctx = context.TODO()

	credential := options.Credential{
		Username: "admin",
		Password: "password",
	}

	clientOptions := options.Client().ApplyURI(url).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

    collection := client.Database("database").Collection("jobs")

	return &DB{collection}, nil
}

func (self *DB) Insert(job *Job) (*Job, error) {
    
	log.Println("create mongodb job entry")

	job.ID = primitive.NewObjectID()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	job.Status = "Submitted"

	_, err := self.Collection.InsertOne(context.Background(), job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (self *DB) GetAll() (*[]Job, error) {

	log.Println("get list of jobs")

	var jobs []Job
	cursor, err := self.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var job Job
		_ = cursor.Decode(&job)
		jobs = append(jobs, job)
	}

	return &jobs, nil
}
