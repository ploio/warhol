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
	"sync"

	"github.com/garyburd/redigo/redis"
)

// RedisClient defines the pub-sub Redis client
type RedisClient struct {
	Conn       redis.Conn
	PubSubConn redis.PubSubConn
	Mutex      sync.Mutex
	MsgCh      chan *Message
}

// NewRedisClient returns a new Redis client
func NewRedisClient(host string, msgChan chan *Message) (*RedisClient, error) {
	log.Printf("[INFO] [redis] PubSub: %s", host)
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:6379", host))
	if err != nil {
		return nil, err
	}
	psc := redis.PubSubConn{Conn: conn}
	//defer conn.Close()
	return &RedisClient{
		Conn:       conn,
		PubSubConn: psc,
		MsgCh:      msgChan,
	}, nil
}

func (client *RedisClient) Subscribe(channels ...interface{}) error {
	for _, channel := range channels {
		log.Printf("[INFO] [redis] Subscribe to [%s]\n", channel)
		client.PubSubConn.Subscribe(channel)
	}
	return nil
}

func (client *RedisClient) Unsubscribe(channels ...interface{}) error {
	for _, channel := range channels {
		log.Printf("[INFO] [redis] Unsubscribe to [%s]\n", channel)
		client.PubSubConn.Unsubscribe(channel)
	}
	return nil
}

func (client *RedisClient) Publish(channel string, message string) {
	log.Printf("[INFO] [redis] Publish: %s to [%s]\n", message, channel)
	client.Mutex.Lock()
	client.Conn.Do("PUBLISH", channel, message)
	client.Mutex.Unlock()
}

func (client *RedisClient) Receive() {
	for {
		switch message := client.PubSubConn.Receive().(type) {
		case redis.Message:
			msg := &Message{
				Type:    "message",
				Channel: message.Channel,
				Data:    string(message.Data),
			}
			log.Printf("[INFO] [redis] Receive: %v", msg)
			client.MsgCh <- msg
		case redis.PMessage:
			msg := &Message{
				Type:    "pmessage",
				Channel: message.Channel,
				Data:    string(message.Data),
			}
			log.Printf("[INFO] [redis] Receive: %v", msg)
			client.MsgCh <- msg
		case redis.Pong:
			msg := &Message{
				Type:    "pong",
				Channel: "",
				Data:    string(message.Data),
			}
			log.Printf("[DEBUG] [redis] Receive: %v", msg)
		case redis.Subscription:
			msg := &Message{
				Type:    message.Kind,
				Channel: message.Channel,
				Data:    string(message.Count),
			}
			log.Printf("[INFO] [redis] Receive: %v", msg)
		default:
			log.Printf("[WARN] [redis] Receive invalid message: %v",
				message)
		}
	}
}
