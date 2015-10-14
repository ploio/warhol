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

	"github.com/labstack/echo"

	"github.com/portefaix/warhol/providers/bitbucket"
)

// BitbucketPushHandler receive a Bitbucket push event notification
func (ws *WebService) BitbucketPushHandler(c *echo.Context) error {
	log.Printf("[INFO] [bitbucket] receive Push event notification")
	var hook bitbucket.PushWebHook
	err := c.Bind(&hook)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&ErrorResponse{
				Error: fmt.Sprintf("Invalid JSON content : %v", err)})
	}
	log.Printf("[DEBUG] [bitbucket] Push webhook: %#v", hook)
	return c.JSON(http.StatusOK, &StatusResponse{Status: "ok"})
}
