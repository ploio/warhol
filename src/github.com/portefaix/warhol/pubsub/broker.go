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
)

const (
	// REDIS PubSub
	REDIS string = "redis"

	// ZEROMQ PubSub
	ZEROMQ string = "zeromq"

	// Channel is for publishers
	Channel = "warhol"
)

var (
	// ErrNotSupported is thrown when the broker is not supported
	ErrNotSupported = errors.New("Broker not supported.")

	// Messaging is the messaging system
	Messaging Broker
)

// Message is a pub-sub message.
type Message struct {
	Type    string
	Channel string
	Data    string
}

// Config represents the broker configuration
type Config struct {
	Type string
	Host string
}

// Broker interface for pubsub clients.
type Broker interface {
	Subscribe(channels ...interface{}) (err error)

	Unsubscribe(channels ...interface{}) (err error)

	Publish(channel string, message string)

	Receive()
}

// NewBroker creates an instance of broker according to the configuration
func NewBroker(config *Config, msgChan chan *Message) (Broker, error) {
	switch config.Type {
	case REDIS:
		return NewRedisClient(config.Host, msgChan)
	case ZEROMQ:
		return NewZeroMQClient(config.Host, msgChan)
	default:
		return nil, ErrNotSupported
	}
}
