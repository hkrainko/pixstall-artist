package rabbitmq

import (
	"bytes"
	"context"
	"github.com/streadway/amqp"
	"log"
	"pixstall-artist/app/domain/artist"
	"pixstall-artist/app/domain/artist/model"
	"time"
)

type ArtistMessageBroker struct {
	artistUseCase artist.UseCase
	ch            *amqp.Channel
}

func NewRabbitMQArtistMessageBroker(useCase artist.UseCase, conn *amqp.Connection) ArtistMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return ArtistMessageBroker{
		artistUseCase: useCase,
		ch:            ch,
	}
}

func (a ArtistMessageBroker) StartArtistQueue(ctx context.Context) {
	//TODO
	q, err := a.ch.QueueDeclare(
		"pixstall-artist_user_artist", // name
		true,                          // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = a.ch.QueueBind(
		q.Name,
		"user.artist.#",
		"user",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue %v", err)
	}

	msgs, err := a.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)

			err := a.artistUseCase.RegisterNewArtist(ctx, &model.Artist{
				ArtistID:         "",
				UserID:           "",
				UserName:         "",
				Email:            "",
				Birthday:         "",
				Gender:           "",
				PhotoURL:         "",
				State:            "",
				FansIDs:          nil,
				LikeIDs:          nil,
				RegistrationTime: time.Time{},
				ArtistIntro:      model.ArtistIntro{},
				ArtistDetails:    model.ArtistDetails{},
				OpenCommissions:  nil,
				Artworks:         nil,
			})
			if err != nil {

			}
		}
	}()

	<-forever
}

func (a ArtistMessageBroker) StopArtistQueue(ctx context.Context) {
	//TODO need?
}
