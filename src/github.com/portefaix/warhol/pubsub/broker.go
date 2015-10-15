// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pubsub

import (
	"errors"
	"fmt"
)

const (
	// REDIS PubSub
	REDIS string = "redis"

	// ZEROMQ PubSub
	ZEROMQ string = "zeromq"
)

var (
	// ErrNotSupported is thrown when the broker is not supported
	ErrNotSupported = errors.New("Broker not supported.")
)

// Message is a pub-sub message.
type Message struct {
	Type    string
	Channel string
	Data    string
}

// Broker interface for pubsub clients.
type Broker interface {
	Subscribe(channels ...interface{}) (err error)

	Unsubscribe(channels ...interface{}) (err error)

	Publish(channel string, message string)

	Receive() (message *Message, err error)
}

// InitBroker creates an instance of Broker
func InitBroker(broker string, host string) (Broker, error) {
	switch broker {
	case REDIS:
		return NewRedisClient(host)
	case ZEROMQ:
		return NewZeroMQClient(host), nil
	default:
		return nil, fmt.Errorf("%s %s", ErrNotSupported.Error(), "")
	}
}
