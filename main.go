package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"pixstall-artist/app/middleware"
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
	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer rabbitmqConn.Close()
	ch, err := rabbitmqConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel %v", err)
	}
	err = ch.ExchangeDeclare(
		"artist",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}
	err = ch.ExchangeDeclare(
		"open-comm",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}

	//gRPC - File
	fileGRPCConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer fileGRPCConn.Close()

	artistMsgBroker := InitArtistMessageBroker(db, rabbitmqConn, fileGRPCConn)
	go artistMsgBroker.StartArtistQueue()
	defer artistMsgBroker.StopArtistQueue()

	commMsgBroker := InitCommissionMessageBroker(db, rabbitmqConn)
	go commMsgBroker.StartValidateQueue()
	defer commMsgBroker.StopAllQueue()

	//Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard: true,
		AllowFiles: true,
		MaxAge: 12 * time.Hour,
	}))

	userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	apiGroup := r.Group("/api")

	artistGroup := apiGroup.Group("/artists")
	{
		ctrl := InitArtistController(db, fileGRPCConn, rabbitmqConn)
		// Artist
		artistGroup.GET("", ctrl.GetArtists)
		artistGroup.GET("/:id", ctrl.GetArtist)
		artistGroup.GET("/:id/details", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetArtistDetails)
		artistGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.UpdateArtist)
		// Open Commission
		artistGroup.GET("/:id/open-commissions", ctrl.GetOpenCommissionsForArtist)
		artistGroup.GET("/:id/open-commissions/details", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetOpenCommissionsDetailsForArtist)
		artistGroup.POST("/:id/open-commissions", userIDExtractor.ExtractPayloadsFromJWT, ctrl.AddOpenCommissionForArtist)
	}

	openCommGroup := apiGroup.Group("/open-commissions")
	{
		ctrl := InitOpenCommissionController(db, fileGRPCConn, rabbitmqConn)
		openCommGroup.GET("/:id", ctrl.GetOpenCommission)
		openCommGroup.GET("", ctrl.GetOpenCommissions)
		openCommGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.UpdateOpenCommission)
		openCommGroup.DELETE("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.DeleteOpenCommission)
	}

	bookmarkGroup := apiGroup.Group("/artists-bookmarks")
	{
		ctrl := InitBookmarkController(db, rabbitmqConn)
		bookmarkGroup.GET("", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetBookmarks)
		bookmarkGroup.GET("/ids", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetBookmarks)
		bookmarkGroup.POST("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.AddBookmark)
		bookmarkGroup.DELETE("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.DeleteBookmark)
	}

	err = r.Run(":9002")
	print(err)
}
