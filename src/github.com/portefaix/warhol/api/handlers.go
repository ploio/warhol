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
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// StatusResponse represents a status response of the REST API
type StatusResponse struct {
	Status string `json:"status"`
}

// ErrorResponse represents an error
type ErrorResponse struct {
	Error string `json:"error"`
}

// WebService represents the Restful API
type WebService struct {
}

// NewWebService creates a new WebService instance
func NewWebService() *WebService {
	log.Debugf("Creates webservice")
	return &WebService{}
}

// Help send a message in JSON
func (ws *WebService) Help(c *echo.Context) error {
	return c.String(http.StatusOK,
		"Welcome to Warhol, the Docker image factory.\n")
}

// GetWebService return a new web service
func GetWebService() *echo.Echo {
	ws := NewWebService()
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.Get("/", ws.Help)
	e.Get("/api/version", ws.DisplayAPIVersion)
	v1 := e.Group("/api/v1")
	v1.Post("/notification/gitlab/push", ws.GitlabPushHandler)
	v1.Post("/notification/gitlab/tag", ws.GitlabTagHandler)
	v1.Post("/notification/gitlab/issue", ws.GitlabIssueHandler)
	return e
}
