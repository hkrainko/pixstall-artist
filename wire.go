//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	artist_deliv_http "pixstall-artist/app/artist/delivery/http"
	artist_deliv_rabbitmq "pixstall-artist/app/artist/delivery/rabbitmq"
	artist_repo "pixstall-artist/app/artist/repo/mongo"
	artist_ucase "pixstall-artist/app/artist/usecase"
	comm_deliv_rabbitmq "pixstall-artist/app/commission/delivery/rabbitmq"
	comm_ucase "pixstall-artist/app/commission/usecase"
	file_repo "pixstall-artist/app/file/repo"
	msg_broker_repo "pixstall-artist/app/msg-broker/repo/rabbitmq"
	opencomm_deliv_http "pixstall-artist/app/open-commission/delivery/http"
	opencomm_repo "pixstall-artist/app/open-commission/repo/mongo"
	opencomm_ucase "pixstall-artist/app/open-commission/usecase"
)

func InitArtistController(db *mongo.Database, grpcConn *grpc.ClientConn, conn *amqp.Connection) artist_deliv_http.ArtistController {
	wire.Build(
		artist_deliv_http.NewArtistController,
		artist_ucase.NewArtistUseCase,
		artist_repo.NewMongoArtistRepo,
		opencomm_ucase.NewOpenCommissionUseCase,
		opencomm_repo.NewMongoOpenCommissionRepo,
		file_repo.NewGRPCFileRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
	)
	return artist_deliv_http.ArtistController{}
}

func InitOpenCommissionController(db *mongo.Database, grpcConn *grpc.ClientConn, conn *amqp.Connection) opencomm_deliv_http.OpenCommissionController {
	wire.Build(
		opencomm_deliv_http.NewOpenCommissionController,
		opencomm_ucase.NewOpenCommissionUseCase,
		opencomm_repo.NewMongoOpenCommissionRepo,
		file_repo.NewGRPCFileRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
	)
	return opencomm_deliv_http.OpenCommissionController{}
}

func InitArtistMessageBroker(db *mongo.Database, conn *amqp.Connection, grpcConn *grpc.ClientConn) artist_deliv_rabbitmq.ArtistMessageBroker {
	wire.Build(
		artist_deliv_rabbitmq.NewRabbitMQArtistMessageBroker,
		artist_ucase.NewArtistUseCase,
		artist_repo.NewMongoArtistRepo,
		opencomm_repo.NewMongoOpenCommissionRepo,
		file_repo.NewGRPCFileRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
	)
	return artist_deliv_rabbitmq.ArtistMessageBroker{}
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection) comm_deliv_rabbitmq.CommissionMessageBroker {
	wire.Build(
		comm_deliv_rabbitmq.NewRabbitMQCommissionMessageBroker,
		comm_ucase.NewCommissionUseCase,
		opencomm_repo.NewMongoOpenCommissionRepo,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		)
	return comm_deliv_rabbitmq.CommissionMessageBroker{}
}
