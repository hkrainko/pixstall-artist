package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"pixstall-artist/app/middleware"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//AWS s3
	awsAccessKey := "AKIA5BWICLKRWX6ARSEF"
	awsSecret := "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"
	token := ""
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(endpoints.ApEast1RegionID),
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   creds,
		},
		//Profile:                 "default", //[default], use [prod], [uat]
		//SharedConfigState:       session.SharedConfigEnable,
	}))
	awsS3 := s3.New(sess)

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
	artistMsgBroker := InitArtistMessageBroker(db, conn, awsS3)
	go artistMsgBroker.StartArtistQueue()
	defer artistMsgBroker.StopArtistQueue()

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
		ctrl := InitArtistController(db, awsS3)
		// Artist
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
		ctrl := InitOpenCommissionController(db)
		openCommGroup.GET("/:id", ctrl.GetOpenCommission)
		openCommGroup.GET("", ctrl.GetOpenCommissions)
		openCommGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.UpdateOpenCommission)
		openCommGroup.DELETE("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.DeleteOpenCommission)
	}

	commissionGroup := apiGroup.Group("/commission")
	{
		ctrl :=
	}

	err = r.Run(":9002")
	print(err)
}
