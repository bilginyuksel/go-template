package broker

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

// Subscriber is a function that handles a message.
// It is used by the broker to handle messages.
// When a message is published, the broker calls all subscribers.
type Subscriber func(context.Context, []byte) error

// Broker is a message broker.
// It is used to publish messages to subscribers.
type Broker struct {
	mutex  *sync.Mutex
	events map[string][]Subscriber
}

// NewBroker creates a new broker.
func New() *Broker {
	return &Broker{
		mutex:  &sync.Mutex{},
		events: make(map[string][]Subscriber),
	}
}

// Publish publishes a message to subscribers.
func (b *Broker) Publish(ctx context.Context, event string, message []byte) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	zap.L().Debug("publishing message to subscribers", zap.String("event", event), zap.ByteString("message", message))

	subscribers := b.events[event]
	for _, subscriber := range subscribers {
		subscriber(ctx, message)
	}
}

// Subscriber adds a subscriber to the broker.
func (b *Broker) Subscribe(event string, handler func(context.Context, []byte) error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	zap.L().Debug("subscribing to event", zap.String("event", event))

	if events, ok := b.events[event]; ok {
		b.events[event] = append(events, handler)
	} else {
		b.events[event] = []Subscriber{handler}
	}
}
