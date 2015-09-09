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
	"fmt"
	"log"
	"net/http"

	//log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/portefaix/warhol/providers/gitlab"
)

// GitlabPushHandler receive a Gitlab push event notification
func (ws *WebService) GitlabPushHandler(c *echo.Context) error {
	log.Printf("[INFO] [gitlab] receive Push event notification")
	var hook gitlab.PushWebhook
	err := c.Bind(&hook)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&ErrorResponse{
				Error: fmt.Sprintf("Invalid JSON content : %v", err)})
	}
	log.Printf("[DEBUG] [gitlab] Push webhook: %#v", hook)
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}

// GitlabTagHandler receive a Gitlab tag event notification
func (ws *WebService) GitlabTagHandler(c *echo.Context) error {
	log.Printf("[INFO] [gitlab] receive Tag event notification")
	var hook gitlab.TagWebhook
	err := c.Bind(&hook)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&ErrorResponse{
				Error: fmt.Sprintf("Invalid JSON content : %v", err)})
	}
	log.Printf("[DEBUG] [gitlab] Tag webhook: %#v", hook)
	log.Printf("[INFO] [gitlab] Tag for project %v", hook.Repository.Name)
	project := ws.Builder.NewProject(
		hook.Repository.Name, "Dockerfile", hook.Repository.URL)
	err = ws.Builder.ToPipeline(project)
	// go ws.Builder.BuildImage(
	// 	hook.Repository.Name,
	// 	//hook.Repository.URL,
	// 	"github.com/nlamirault/aneto",
	// 	"Dockerfile")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&ErrorResponse{
				Error: fmt.Sprintf("Can't manage hook : %v", err)})
	}
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}

// GitlabIssueHandler receive a Gitlab tag event notification
func (ws *WebService) GitlabIssueHandler(c *echo.Context) error {
	log.Printf("[INFO] [gitlab] receive Issue event notification")
	var hook gitlab.IssueWebhook
	err := c.Bind(&hook)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&ErrorResponse{
				Error: fmt.Sprintf("Invalid JSON content : %v", err)})
	}
	log.Printf("[INFO] [gitlab] Issue webhook: %#v", hook)
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}
