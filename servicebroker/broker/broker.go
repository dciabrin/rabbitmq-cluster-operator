package broker

import (
	"context"
	"errors"

	"github.com/pivotal-cf/brokerapi"
)

type RabbitMQServiceBroker struct {
	Config Config
}

func New(cfg Config) brokerapi.ServiceBroker {
	return RabbitMQServiceBroker{
		Config: cfg,
	}
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) GetInstance(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {
	return brokerapi.GetInstanceDetailsSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) GetBinding(ctx context.Context, instanceID, bindingID string) (brokerapi.GetBindingSpec, error) {
	return brokerapi.GetBindingSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("Not implemented")
}