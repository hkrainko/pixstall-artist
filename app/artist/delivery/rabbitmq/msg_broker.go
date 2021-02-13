package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	reg_info "pixstall-artist/app/artist/delivery/model/reg-artist"
	update_artist "pixstall-artist/app/artist/delivery/model/update-artist"
	"pixstall-artist/domain/artist"
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

func (a ArtistMessageBroker) StartArtistQueue() {
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
		"user.#",
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
			d.Ack(false)

			ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
			go func() {
				for {
					select {
					case <-ctx.Done():
						switch ctx.Err() {
						case context.DeadlineExceeded:
							log.Println("context.DeadlineExceeded")
						case context.Canceled:
							log.Println("context.Canceled")
						default:
							log.Println("default")
						}
						return // returning not to leak the goroutine
					}
				}
			}()

			switch d.RoutingKey {
			case "user.new.isArtist":
				err := a.registerNewArtist(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			case "user.update.isArtist":
				err := a.UpdateArtistUser(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			default:
				cancel()
			}


		}
	}()

	<-forever
}

func (a ArtistMessageBroker) StopArtistQueue() {
	err := a.ch.Close()
	if err != nil {
		log.Printf("StopArtistQueue err %v", err)
	}
	log.Printf("StopArtistQueue success")
}

func (a ArtistMessageBroker) registerNewArtist(ctx context.Context, body []byte) error {
	req := reg_info.Request{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return a.artistUseCase.RegisterNewArtist(ctx, req.RegInfo)
}

func (a ArtistMessageBroker) UpdateArtistUser(ctx context.Context, body []byte) error {
	req := update_artist.Request{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	id, err := a.artistUseCase.UpdateArtistUser(ctx, *req.UserUpdater)
	if err != nil {
		return err
	}
	log.Printf("UpdateArtistUser success ID: %v", id)
	return nil
}
