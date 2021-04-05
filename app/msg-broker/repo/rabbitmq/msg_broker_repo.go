package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	model2 "pixstall-artist/app/msg-broker/repo/rabbitmq/msg"
	model4 "pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/commission/model"
	msg_broker "pixstall-artist/domain/msg-broker"
	model3 "pixstall-artist/domain/open-commission/model"
)

type rabbitmqMsgBrokerRepo struct {
	ch *amqp.Channel
}

func NewRabbitMQMsgBrokerRepo(conn *amqp.Connection) msg_broker.Repo {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}
	return rabbitmqMsgBrokerRepo{
		ch: ch,
	}
}

func (r rabbitmqMsgBrokerRepo) SendCommOpenCommValidationMsg(ctx context.Context, validation model.CommissionOpenCommissionValidation) error {
	vComm := model2.CommissionOpenCommissionValidation{
		CommissionOpenCommissionValidation: validation,
	}
	b, err := json.Marshal(vComm)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"commission",
		"commission.event.validation.open-comm",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r rabbitmqMsgBrokerRepo) SendArtistCreatedMsg(ctx context.Context, artist model4.Artist) error {
	createdArtist := model2.NewCreatedArtist(artist)
	b, err := json.Marshal(createdArtist)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"artist",
		"artist.event.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r rabbitmqMsgBrokerRepo) SendArtistUpdatedMsg(ctx context.Context, updater model4.ArtistUpdater) error {
	updatedArtist := model2.NewUpdatedArtist(updater)
	b, err := json.Marshal(updatedArtist)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"artist",
		"artist.event.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r rabbitmqMsgBrokerRepo) SendOpenCommCreatedMsg(ctx context.Context, openComm model3.OpenCommission) error {
	createdOpenComm := model2.NewCreatedOpenCommission(openComm)
	b, err := json.Marshal(createdOpenComm)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"open-comm",
		"open-comm.event.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r rabbitmqMsgBrokerRepo) SendOpenCommUpdatedMsg(ctx context.Context, openComm model3.OpenCommission) error {
	return nil
}
