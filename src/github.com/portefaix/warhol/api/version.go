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
	// "fmt"
	//"log"
	"net/http"

	//log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// Version represents version of the REST API
type Version struct {
	Version string `json:"version"`
}

// DisplayAPIVersion sends the API version in JSON format
func (ws *WebService) DisplayAPIVersion(c *echo.Context) error {
	return c.JSON(http.StatusOK, &Version{Version: "1"})
}
