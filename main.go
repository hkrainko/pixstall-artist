package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//Mongo
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer cancel()
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := dbClient.Database("pixstall-artist")

	//RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer conn.Close()
	artistMsgBroker := InitArtistMessageBroker(db, conn)
	go artistMsgBroker.StartArtistQueue(ctx)
	defer artistMsgBroker.StopArtistQueue(ctx)

	//Gin
	r := gin.Default()

	authGroup := r.Group("/artist")
	{
		ctr := InitArtistController(db)
		authGroup.POST("/getArtist", ctr.GetArtist)
		authGroup.POST("/updateArtist", ctr.UpdateArtist)
	}

	err = r.Run(":9002")
	print(err)
}