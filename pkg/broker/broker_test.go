package broker_test

import (
	"context"
	"errors"
	"testing"

	"gotemplate/pkg/broker"

	"github.com/stretchr/testify/assert"
)

type testSubscriber struct {
	count    int
	messages [][]byte
}

func (t *testSubscriber) Handle(ctx context.Context, message []byte) error {
	t.count++
	t.messages = append(t.messages, message)

	if string(message) == "msg1" {
		return errors.New("something went wrong")
	}

	return nil
}

func TestBroker(t *testing.T) {
	b := broker.New()
	t1 := testSubscriber{}
	t2 := testSubscriber{}

	b.Subscribe("e1", t1.Handle)
	b.Subscribe("e1", t2.Handle)

	msgs := [][]byte{
		[]byte("msg1"),
		[]byte("msg2"),
		[]byte("msg3"),
	}

	b.Publish(context.TODO(), "e1", msgs[0])
	b.Publish(context.TODO(), "e1", msgs[1])
	b.Publish(context.TODO(), "e1", msgs[2])

	b.Publish(context.TODO(), "e1", msgs[0])
	b.Publish(context.TODO(), "e1", msgs[1])

	assert.Equal(t, 5, t1.count)
	assert.Equal(t, 5, t2.count)
	assert.Equal(t, append(msgs, msgs[0], msgs[1]), t1.messages)
	assert.Equal(t, append(msgs, msgs[0], msgs[1]), t2.messages)
}
