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
	"time"

	zmq "github.com/pebbe/zmq4"
)

// ZeroMQClient - just defines the pub and sub ZMQ sockets.
type ZeroMQClient struct {
	Pub   *zmq.Socket
	Sub   *zmq.Socket
	MsgCh chan *Message
}

// NewZeroMQClient creates a new ZeroMQClient instance
func NewZeroMQClient(host string, msgChan chan *Message) (*ZeroMQClient, error) {
	log.Printf("[INFO] [zeromq] PubSub: %s", host)
	pub, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		return nil, err
	}
	pub.Bind(fmt.Sprintf("tcp://%s:%d", host, 5556))
	//  Ensure subscriber connection has time to complete
	time.Sleep(time.Second)

	sub, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, err
	}
	sub.Connect(fmt.Sprintf("tcp://%s:%d", host, 5556))
	return &ZeroMQClient{
		Pub:   pub,
		Sub:   sub,
		MsgCh: msgChan,
	}, nil
}

// Subscribe establish message filter
func (client *ZeroMQClient) Subscribe(channels ...interface{}) error {
	for _, channel := range channels {
		log.Printf("[INFO] [zeromq] Subscribe to [%s]\n", channel)
		client.Sub.SetSubscribe(channel.(string))
	}
	return nil
}

// Unsubscribe remove message filter
func (client *ZeroMQClient) Unsubscribe(channels ...interface{}) error {
	for _, channel := range channels {
		log.Printf("[INFO] [zeromq] Unsubscribe to [%s]\n", channel)
		client.Sub.SetUnsubscribe(channel.(string))
	}
	return nil
}

// Publish send a message to a channel
func (client *ZeroMQClient) Publish(channel, message string) {
	log.Printf("[INFO] [zeromq] Publish: %s to [%s]\n", message, channel)
	client.Pub.Send(fmt.Sprintf("%s %s", channel, message), 0)
}

// Receive receive  parts as message from socket.
func (client *ZeroMQClient) Receive() {
	for {
		message, err := client.Sub.RecvMessage(0)
		log.Printf("[INFO] [zeromq] Receive: %s \n", message)
		if err != nil {
			log.Printf("[WARN] [zeromq] Error receiving message: %s",
				err.Error())
		}
		topic := message[0]
		data := message[1]
		if topic != Channel {
			log.Printf("[WARN] [zeromq] Invalid message: %s %s",
				topic, data)
		}
		msg := &Message{
			Type:    "message",
			Channel: topic,
			Data:    data,
		}
		log.Printf("[INFO] [zeromq] Message: %v", msg)
	}
}
