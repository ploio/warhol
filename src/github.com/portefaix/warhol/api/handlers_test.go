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

package api

import (
	//"fmt"
	"testing"

	//"github.com/labstack/echo"

	"github.com/portefaix/warhol/providers/docker"
	"github.com/portefaix/warhol/pubsub"
)

var api = map[string]string{
	"/":                                 "GET",
	"/api/version":                      "GET",
	"/api/v1/notification/gitlab/push":  "POST",
	"/api/v1/notification/gitlab/tag":   "POST",
	"/api/v1/notification/gitlab/issue": "POST",
}

func Test_WebServiceRoutes(t *testing.T) {
	builder, _ := docker.NewBuilder(
		"127.0.0.1",
		false,
		"/tmp",
		"127.0.0.1:5000",
		&docker.Authentication{
			Username: "",
			Password: "",
			Email:    "",
		},
		&pubsub.Config{Type: pubsub.ZEROMQ})
	ws := GetWebService(builder)
	// fmt.Printf("Routes : %#v", ws.Routes())
	routes := ws.Routes()
	if len(routes) != 5 {
		t.Fatalf("Invalid routes. : %v", routes)
	}
	for _, route := range ws.Routes() {
		if api[route.Path] != route.Method {
			t.Fatalf("Unknown route. : %v", route)
		}
	}
}
