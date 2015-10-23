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

package irc

import (
	"log"

	"github.com/fabioxgn/go-bot"

	"github.com/portefaix/warhol/pubsub"
)

type Publisher struct {
	Config  *bot.Config
	Broker  pubsub.Broker
	MsgChan chan *pubsub.Message
}

type Config struct {
	Server   string
	Channel  string
	Username string
	Nickname string
	Password string
}

// NewPublisher creates a new IRC bot publisher
func NewPublisher(config *Config, brokerConf *pubsub.Config, debug bool) (*Publisher, error) {
	msgChan := make(chan *pubsub.Message)
	sub, err := pubsub.NewBroker(brokerConf, msgChan)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		Config: &bot.Config{
			Server:   config.Server,
			Channels: []string{config.Channel},
			User:     config.Username,
			Nick:     config.Nickname,
			Password: config.Password,
			UseTLS:   true,
			Debug:    debug,
		},
		Broker:  sub,
		MsgChan: msgChan,
	}, nil
}

// Run connect to the specified IRC server and starts the bot
func (p *Publisher) Run() {
	log.Printf("[INFO] [irc] IRC Publisher run: %v", p.Config)
	bot.RegisterCommand(
		"version",
		"Display version of Warhol.",
		"",
		versionCmd)
	p.Broker.Subscribe(pubsub.Channel)
	go p.Broker.Receive()
	go p.Serve()
	go bot.Run(p.Config)
}

func (p *Publisher) Serve() {
	for {
		message := <-p.MsgChan
		log.Printf("[INFO] [irc] Message receive : %v", message)
	}
}
