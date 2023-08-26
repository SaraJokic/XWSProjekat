package nats

import (
	saga "accommodationsBackend/common/saga/messaging"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	component *NATSComponent
	conn      *nats.EncodedConn
	subject   string
}

func NewNATSPublisher(component *NATSComponent, subject string) (saga.Publisher, error) {
	conn := component.NATS()
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		component: component,
		conn:      encConn,
		subject:   subject,
	}, nil
}

func (p *Publisher) Publish(message interface{}) error {
	err := p.conn.Publish(p.subject, message)
	if err != nil {
		return err
	}
	return nil
}
