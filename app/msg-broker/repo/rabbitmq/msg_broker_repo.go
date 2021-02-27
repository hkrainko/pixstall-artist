package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	model2 "pixstall-artist/app/msg-broker/repo/rabbitmq/model"
	msg_broker "pixstall-artist/domain/msg-broker"
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

func (r rabbitmqMsgBrokerRepo) SendValidatedCommissionMsg(ctx context.Context, err error) error {
	vComm := model2.ValidatedCommission{
		IsValid: err == nil,
		Reason: getRejectReason(err),
	}
	b, err := json.Marshal(vComm)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"commission",
		"commission.event.validated",
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

// Private
func getRejectReason(err error) *string {
	if err != nil {
		errStr := err.Error()
		return &errStr
	}
	return nil
}
