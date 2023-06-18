package mq

import (
	"ji/config"
	"strconv"
	"time"

	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
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

func (rc *RabbitMQClient) SendMessageDirect(message []byte, exchangeName string, queueName string) error {
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
			Body:         message,
			Timestamp:    time.Now(),
		},
	)
	return err
}

func (rc *RabbitMQClient) SendMessageDelay(message []byte, exchangeName string, queueName string, ttl int64) error {

	dlxExchangeName := exchangeName + "dlx"
	dlxQueueName := queueName + "dlx"
	routekey := queueName
	expireTime := strconv.FormatInt(ttl, 10)

	// 死信队列
	if err := rc.channel.ExchangeDeclare(dlxExchangeName, EXCHANGE_TYPE, true, false, false, false, nil); err != nil {
		return err
	}
	if _, err := rc.channel.QueueDeclare(dlxQueueName, true, false, false, true, nil); err != nil {
		return err
	}
	if err := rc.channel.QueueBind(dlxQueueName, routekey, dlxExchangeName, true, nil); err != nil {
		return err
	}

	// 将死信队列绑定至普通队列
	if err := rc.channel.ExchangeDeclare(exchangeName, EXCHANGE_TYPE, true, false, false, false, nil); err != nil {
		return err
	}
	args := amqp.Table{
		"x-dead-letter-exchange": dlxExchangeName,
	}
	if _, err := rc.channel.QueueDeclare(queueName, true, false, false, true, args); err != nil {
		return err
	}
	if err := rc.channel.QueueBind(queueName, routekey, exchangeName, true, nil); err != nil {
		return err
	}
	// 发送消息
	if err := rc.channel.Publish(exchangeName, routekey, false, false,
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "application/json",
			Body:         message,
			Timestamp:    time.Now(),
			Expiration:   expireTime,
		},
	); err != nil {
		return err
	}

	return nil
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

func (rc *RabbitMQClient) ConsumerDelay(exchangeName string, queueName string, consumerHandler func(delivery amqp.Delivery) error) error {
	dlxQueueName := queueName + "dlx"
	consumerTag := dlxQueueName

	if _, err := rc.channel.QueueDeclare(dlxQueueName, true, false, false, true, nil); err != nil {
		return err
	}

	var delivery <-chan amqp.Delivery
	var err error
	if delivery, err = rc.channel.Consume(dlxQueueName, consumerTag, false, false, false, false, nil); err != nil {
		return err
	}

	go func() {
		for d := range delivery {
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
	}()
	return nil
}

func (rc *RabbitMQClient) Close() {
	rc.channel.Close()
	rc.connection.Close()
}
