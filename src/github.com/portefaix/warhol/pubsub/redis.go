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
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

const (
	// SOCKET represents the Docker socket endpoint
	SOCKET = "unix:///var/run/docker.sock"

	// REGISTRY is the default Docker registry
	REGISTRY = "index.docker.io"

	// ChannelPub is for publishers
	Channel = "warhol"

	// ChannelSub is for subscribers
	//ChannelSub = "warhol-sub"
)

// RedisClient defines the pub-sub Redis client
type RedisClient struct {
	Conn       redis.Conn
	PubSubConn redis.PubSubConn
}

// NewRedisClient returns a new Redis client
func NewRedisClient(host string) (*RedisClient, error) {
	log.Printf("[INFO] [Redis] PubSub: %s", host)
	host = fmt.Sprintf("%s:6379", host)
	conn, err := redis.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	psc := redis.PubSubConn{Conn: conn}
	return &RedisClient{
		Conn:       conn,
		PubSubConn: psc,
	}, nil
}

func (client *RedisClient) Publish(channel string, message string) {
	log.Printf("[INFO] [redis] Publish: %s to %s channel\n", message, channel)
	client.Conn.Send("PUBLISH", channel, message)
}

func (client *RedisClient) Receive() (*Message, error) {
	switch message := client.PubSubConn.Receive().(type) {
	case redis.Message:
		msg := &Message{
			Type:    "message",
			Channel: message.Channel,
			Data:    string(message.Data),
		}
		log.Printf("[INFO] [redis] Receive: %s from %s channel\n",
			msg.Channel, msg.Data)
		return msg, nil
	case redis.Subscription:
		return &Message{
			Type:    message.Kind,
			Channel: message.Channel,
			Data:    string(message.Count),
		}, nil
	}
	return nil, nil
}

func (client *RedisClient) Subscribe(channels ...interface{}) error {
	for _, channel := range channels {
		log.Printf("[INFO] [redis] Subscribe to %s\n", channel)
		client.PubSubConn.Subscribe(channel)
	}
	return nil
}

func (client *RedisClient) Unsubscribe(channels ...interface{}) error {
	for _, channel := range channels {
		client.PubSubConn.Unsubscribe(channel)
	}
	return nil
}
