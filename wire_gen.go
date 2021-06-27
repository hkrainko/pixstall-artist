// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"pixstall-artist/app/artist/delivery/http"
	rabbitmq2 "pixstall-artist/app/artist/delivery/rabbitmq"
	mongo2 "pixstall-artist/app/artist/repo/mongo"
	"pixstall-artist/app/artist/usecase"
	http3 "pixstall-artist/app/bookmark/delivery/http"
	usecase3 "pixstall-artist/app/bookmark/usecase"
	rabbitmq3 "pixstall-artist/app/commission/delivery/rabbitmq"
	usecase4 "pixstall-artist/app/commission/usecase"
	"pixstall-artist/app/file/repo"
	"pixstall-artist/app/msg-broker/repo/rabbitmq"
	http2 "pixstall-artist/app/open-commission/delivery/http"
	mongo3 "pixstall-artist/app/open-commission/repo/mongo"
	usecase2 "pixstall-artist/app/open-commission/usecase"
)

// Injectors from wire.go:

func InitArtistController(db *mongo.Database, grpcConn *grpc.ClientConn, conn *amqp.Connection) http.ArtistController {
	artistRepo := mongo2.NewMongoArtistRepo(db)
	open_commissionRepo := mongo3.NewMongoOpenCommissionRepo(db)
	fileRepo := repo.NewGRPCFileRepository(grpcConn)
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	useCase := usecase.NewArtistUseCase(artistRepo, open_commissionRepo, fileRepo, msg_brokerRepo)
	open_commissionUseCase := usecase2.NewOpenCommissionUseCase(open_commissionRepo, msg_brokerRepo, fileRepo)
	artistController := http.NewArtistController(useCase, open_commissionUseCase)
	return artistController
}

func InitOpenCommissionController(db *mongo.Database, grpcConn *grpc.ClientConn, conn *amqp.Connection) http2.OpenCommissionController {
	open_commissionRepo := mongo3.NewMongoOpenCommissionRepo(db)
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	fileRepo := repo.NewGRPCFileRepository(grpcConn)
	useCase := usecase2.NewOpenCommissionUseCase(open_commissionRepo, msg_brokerRepo, fileRepo)
	openCommissionController := http2.NewOpenCommissionController(useCase)
	return openCommissionController
}

func InitBookmarkController(db *mongo.Database, conn *amqp.Connection) http3.BookmarkController {
	artistRepo := mongo2.NewMongoArtistRepo(db)
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	useCase := usecase3.NewBookmarkUseCase(artistRepo, msg_brokerRepo)
	bookmarkController := http3.NewBookmarkController(useCase)
	return bookmarkController
}

func InitArtistMessageBroker(db *mongo.Database, conn *amqp.Connection, grpcConn *grpc.ClientConn) rabbitmq2.ArtistMessageBroker {
	artistRepo := mongo2.NewMongoArtistRepo(db)
	open_commissionRepo := mongo3.NewMongoOpenCommissionRepo(db)
	fileRepo := repo.NewGRPCFileRepository(grpcConn)
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	useCase := usecase.NewArtistUseCase(artistRepo, open_commissionRepo, fileRepo, msg_brokerRepo)
	artistMessageBroker := rabbitmq2.NewRabbitMQArtistMessageBroker(useCase, conn)
	return artistMessageBroker
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection) rabbitmq3.CommissionMessageBroker {
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	open_commissionRepo := mongo3.NewMongoOpenCommissionRepo(db)
	useCase := usecase4.NewCommissionUseCase(msg_brokerRepo, open_commissionRepo)
	commissionMessageBroker := rabbitmq3.NewRabbitMQCommissionMessageBroker(useCase, conn)
	return commissionMessageBroker
}
