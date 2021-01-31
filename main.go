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
	//Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	go artistMsgBroker.StartArtistQueue()
	defer artistMsgBroker.StopArtistQueue()

	//Gin
	r := gin.Default()

	apiGroup := r.Group("/api")

	artistGroup := apiGroup.Group("/artists")
	{
		ctrl := InitArtistController(db)
		artistGroup.GET("/:id", ctrl.GetArtist)
		artistGroup.PATCH("/:id", ctrl.UpdateArtist)
		artistGroup.GET("/:id/openCommissions", ctrl.GetOpenCommissionsForArtist)
		artistGroup.POST("/:id/openCommissions", ctrl.AddOpenCommissionForArtist)
		artistGroup.PATCH("/:id/openCommissions/:openCommId", ctrl.UpdateOpenCommissionForArtist)
		artistGroup.DELETE("/:id/openCommissions/:openCommId", ctrl.DeleteOpenCommissionForArtist)
	}

	openCommGroup := apiGroup.Group("/open-commissions")
	{
		ctrl := InitOpenCommissionController(db)
		openCommGroup.GET("/:id", ctrl.GetOpenCommission)
		openCommGroup.GET("", ctrl.GetOpenCommissions)
	}

	err = r.Run(":9002")
	print(err)
}
