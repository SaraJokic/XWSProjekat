package nats

import (
	saga "accommodationsBackend/common/saga/messaging"
	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	component  *NATSComponent
	conn       *nats.EncodedConn
	subject    string
	queueGroup string
}

func NewNATSSubscriber(component *NATSComponent, subject, queueGroup string) (saga.Subscriber, error) {
	conn := component.NATS()
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		component:  component,
		conn:       encConn,
		subject:    subject,
		queueGroup: queueGroup,
	}, nil
}

func (s *Subscriber) Subscribe(handler interface{}) error {
	_, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, handler)
	if err != nil {
		return err
	}
	return nil
}
