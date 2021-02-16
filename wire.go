//+build wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	artist_deliv_http "pixstall-artist/app/artist/delivery/http"
	artist_deliv_rabbitmq "pixstall-artist/app/artist/delivery/rabbitmq"
	artist_repo "pixstall-artist/app/artist/repo/mongo"
	artist_ucase "pixstall-artist/app/artist/usecase"
	image_repo "pixstall-artist/app/image/aws-s3"
	opencomm_deliv_http "pixstall-artist/app/open-commission/delivery/http"
	opencomm_repo "pixstall-artist/app/open-commission/repo/mongo"
	opencomm_ucase "pixstall-artist/app/open-commission/usecase"
	commission_deliv_http "pixstall-artist/app/commission/delivery/http"
)

func InitArtistController(db *mongo.Database, awsS3 *s3.S3) artist_deliv_http.ArtistController {
	wire.Build(
		artist_deliv_http.NewArtistController,
		artist_ucase.NewArtistUseCase,
		artist_repo.NewMongoArtistRepo,
		opencomm_repo.NewMongoOpenCommissionRepo,
		image_repo.NewAWSS3ImageRepository,
	)
	return artist_deliv_http.ArtistController{}
}

func InitOpenCommissionController(db *mongo.Database) opencomm_deliv_http.OpenCommissionController {
	wire.Build(
		opencomm_deliv_http.NewOpenCommissionController,
		opencomm_ucase.NewOpenCommissionUseCase,
		opencomm_repo.NewMongoOpenCommissionRepo,
	)
	return opencomm_deliv_http.OpenCommissionController{}
}

func InitArtistMessageBroker(db *mongo.Database, conn *amqp.Connection, awsS3 *s3.S3) artist_deliv_rabbitmq.ArtistMessageBroker {
	wire.Build(
		artist_deliv_rabbitmq.NewRabbitMQArtistMessageBroker,
		artist_ucase.NewArtistUseCase,
		artist_repo.NewMongoArtistRepo,
		opencomm_repo.NewMongoOpenCommissionRepo,
		image_repo.NewAWSS3ImageRepository,
	)
	return artist_deliv_rabbitmq.ArtistMessageBroker{}
}

func InitCommissionController(conn *amqp.Connection) commission_deliv_http.CommissionController {
	wire.Build(

	)
	return commission_deliv_http.CommissionController{}
}
