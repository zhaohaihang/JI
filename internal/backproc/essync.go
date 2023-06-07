package backproc

import (
	"encoding/json"
	"ji/internal/serializer"
	"ji/pkg/es"
	"ji/pkg/mq"
	"strconv"

	"github.com/streadway/amqp"
)

type EsSyncProc struct {
	ec *es.EsClient
	rm *mq.RabbitMQClient
}

func NewEsSyncProc(ec *es.EsClient, rm *mq.RabbitMQClient) *EsSyncProc {
	return &EsSyncProc{
		ec: ec,
		rm: rm,
	}
}

func (esp *EsSyncProc) start() error {
	if err := esp.rm.ConsumerDirect("activityExChange", "activityCreateQueue", esp.activityCreateHandler); err != nil {
		return err
	}

	if err := esp.rm.ConsumerDirect("activityExChange", "activityUpdateQueue", esp.activityUpdateHandler); err != nil {
		return err
	}
	return nil
}

func (esp *EsSyncProc) stop() {
	esp.rm.Close()
}

func (esp *EsSyncProc) activityCreateHandler(delivery amqp.Delivery) error {
	
	var activity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &activity); err != nil {
		return err
	}
	// TODO sync to ES
	Params := make(map[string]string)
	Params["index"] ="activity"
	Params["id"] = strconv.Itoa(int(activity.ID))
	Params["bodyJson"] = string(delivery.Body)
	esp.ec.Create(Params)
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
