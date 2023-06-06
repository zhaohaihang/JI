package mq

import (
	"ji/config"
	"time"

	"github.com/google/wire"
	"github.com/streadway/amqp"
)

const (
	EXCHANGE_TYPE = "direct"
)

type RabbitMQClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQClient(config *config.Config) (*RabbitMQClient, error) {
	var rabbitMQClient RabbitMQClient
	var err error

	url := "amqp://" + config.RabbitMQ.Username + ":" + config.RabbitMQ.Password + "@" + config.RabbitMQ.Host + ":" + config.RabbitMQ.Port + "/"

	if rabbitMQClient.connection, err = amqp.Dial(url); err != nil {
		return nil, err
	}
	if rabbitMQClient.channel, err = rabbitMQClient.connection.Channel(); err != nil {
		rabbitMQClient.connection.Close()
		return nil, err
	}

	return &rabbitMQClient, nil
}

var RabbitMQClientProviderSet = wire.NewSet(NewRabbitMQClient)

func (rc *RabbitMQClient) SendMessageDirect(message string, exchangeName string, queueName string) error {
	routekey := queueName
	if err := rc.channel.ExchangeDeclare(exchangeName, EXCHANGE_TYPE, true, false, false, false, nil); err != nil {
		return err
	}

	if _, err := rc.channel.QueueDeclare(queueName, true, false, false, true, nil); err != nil {
		return err
	}

	if err := rc.channel.QueueBind(queueName, routekey, exchangeName, true, nil); err != nil {
		return err
	}
	err := rc.channel.Publish(exchangeName, routekey, false, false,
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "application/json",
			Body:         []byte(message),
			Timestamp:    time.Now(),
		},
	)
	return err
}

func (rc *RabbitMQClient) ConsumerDirect(exchangeName string, queueName string, consumerHandler func(delivery amqp.Delivery) error) error {
	routekey := queueName
	consumerTag := queueName
	if err := rc.channel.ExchangeDeclare(exchangeName, EXCHANGE_TYPE, true, false, false, false, nil); err != nil {
		return err
	}

	if _, err := rc.channel.QueueDeclare(queueName, true, false, false, true, nil); err != nil {
		return err
	}

	if err := rc.channel.QueueBind(queueName, routekey, exchangeName, true, nil); err != nil {
		return err
	}
	var delivery <-chan amqp.Delivery
	var err error
	if delivery, err = rc.channel.Consume(queueName, consumerTag, false, false, false, false, nil); err != nil {
		return err
	}

	go func() {
		for {
			select {
			case d := <-delivery:
				if consumerHandler == nil {
					_ = d.Ack(false)
					continue
				}
				if err := consumerHandler(d); err == nil {
					_ = d.Ack(false)
				} else {
					_ = d.Nack(false, false)
				}
			}
		}
	}()
	return nil
}
