//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	artist_deliv_http "pixstall-artist/app/artist/delivery/http"
	artist_deliv_rabbitmq "pixstall-artist/app/artist/delivery/rabbitmq"
	artist_repo "pixstall-artist/app/artist/repo/mongo"
	artist_ucase "pixstall-artist/app/artist/usecase"
)

func InitArtistController(db *mongo.Database) artist_deliv_http.ArtistController {
	wire.Build(
		artist_deliv_http.NewArtistController,
		artist_ucase.NewArtistUseCase,
		artist_repo.NewMongoArtistRepo,
	)
	return artist_deliv_http.ArtistController{}
}

func InitArtistMessageBroker(db *mongo.Database, conn *amqp.Connection) artist_deliv_rabbitmq.ArtistMessageBroker {
	wire.Build(
		artist_deliv_rabbitmq.NewRabbitMQArtistMessageBroker,
		artist_ucase.NewArtistUseCase,
		artist_repo.NewMongoArtistRepo,
	)
	return artist_deliv_rabbitmq.ArtistMessageBroker{}
}
