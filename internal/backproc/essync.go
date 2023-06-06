package backproc

import (
	"encoding/json"
	"ji/pkg/es"
	"ji/pkg/mq"

	"github.com/streadway/amqp"
)

type EsSyncProc struct {
	ec *es.EsClient
	rm *mq.RabbitMQClient
}

func NewEsSyncProc() *EsSyncProc {
	var esSyncProc EsSyncProc
	return &esSyncProc
}

func (esp *EsSyncProc) Start() {
	esp.rm.ConsumerDirect("activityExChange", "activityInsertQueue", esp.activityNewHandler)
	esp.rm.ConsumerDirect("activityExChange", "activityUpdateQueue", esp.activityUpdateHandler)
}

func (esp *EsSyncProc) activityNewHandler(delivery amqp.Delivery) error {
	var mqData map[string]interface{}
	if err := json.Unmarshal(delivery.Body, &mqData); err != nil {
		return err
	}
	// TODO sync to ES
	return nil
}

func (esp *EsSyncProc) activityUpdateHandler(delivery amqp.Delivery) error {
	var mqData map[string]interface{}
	if err := json.Unmarshal(delivery.Body, &mqData); err != nil {
		return err
	}
	// TODO sync to ES
	return nil
}
