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
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/portefaix/warhol/providers/gitlab"
)

// GitlabPushHandler receive a Gitlab push event notification
func (ws *WebService) GitlabPushHandler(c *echo.Context) error {
	log.Infof("Gitlab receive Push event notification")
	var hook gitlab.PushWebhook
	c.Bind(&hook)
	log.Debugf("Gitlab Push webhook: %s", hook)
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}

// GitlabTagHandler receive a Gitlab tag event notification
func (ws *WebService) GitlabTagHandler(c *echo.Context) error {
	log.Infof("Gitlab receive Tag event notification")
	var hook gitlab.TagWebhook
	c.Bind(&hook)
	log.Debugf("Gitlab Tag webhook: %s", hook)
	log.Infof("Gitlab Tag for project %s", hook.ProjectID)
	//bot := workers.Workers["irc"]
	//go bot.NotifyTag(hook)
	// ipp := workers.Workers["ipp"]
	// go ipp.NotifyTag(hook)
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}

// GitlabIssueHandler receive a Gitlab tag event notification
func (ws *WebService) GitlabIssueHandler(c *echo.Context) error {
	log.Infof("Gitlab receive Issue event notification")
	var hook gitlab.IssueWebhook
	c.Bind(&hook)
	log.Debugf("Gitlab Issue webhook: %s", hook)
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}
