package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"log"
	"pixstall-artist/domain/commission/model"
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

func (r rabbitmqMsgBrokerRepo) SendAddCommissionMsg(ctx context.Context, creator model.CommissionCreator) error {
	panic("implement me")
}
