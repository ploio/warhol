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
	//"fmt"
	"testing"
)

func Test_UnsupportedBroker(t *testing.T) {
	broker, err := NewBroker(&Config{
		Type: "Foo",
		Host: "localhost",
	}, nil)
	if broker != nil || err == nil {
		t.Fatalf("Error retrieve an unsupported broker")
	}
}

func Test_RedisBroker(t *testing.T) {
	broker, err := NewBroker(&Config{
		Type: REDIS,
		Host: "127.0.0.1",
	}, nil)
	if broker == nil || err != nil {
		t.Fatalf("Error creating Redis broker")
	}
}

func Test_ZeroMQBroker(t *testing.T) {
	broker, err := NewBroker(&Config{
		Type: ZEROMQ,
	}, nil)
	if broker == nil || err != nil {
		t.Fatalf("Error creating ZeroMQ broker")
	}
}
