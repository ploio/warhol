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

package gitlab

import (
	"strings"
)

// Repository represents the repository content from webhook
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Home        string `json:"home"`
	GitHTTPURL  string `json:"git_http_url"`
	GitSSHURL   string `json:"git_ssh_url"`
}

// Commit represents commit information from the webhook
type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

// Author represents author information from the webhook
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// User represents the user information from the webhook
type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// Attributes represents the object attributes from the webhook
type Attributes struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	AssigneeID  int    `json:"assignee_id"`
	AuthorID    int    `json:"author_id"`
	ProjectID   int    `json:"project_id"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Position    int    `json:"position"`
	BranchName  string `json:"branch_name"`
	Description string `json:"description"`
	MilestoneID int    `json:"milestone_id"`
	State       string `json:"state"`
	IID         int    `json:"iid"`
	URL         string `json:"url"`
	Action      string `json:"action"`
}

// PushWebhook represents push information from the webhook
type PushWebhook struct {
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	Username          string     `json:"user_name"`
	UserID            int        `json:"user_id"`
	ProjectID         int        `json:"project_id"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int        `json:"total_commits_count"`
}

// TagWebhook represents tag information from the webhook
type TagWebhook struct {
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	Username          string     `json:"user_name"`
	UserID            int        `json:"user_id"`
	ProjectID         int        `json:"project_id"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int        `json:"total_commits_count"`
}

// IssueWebhook represents issue information from the webhook
type IssueWebhook struct {
	Kind       string     `json:"object_kind"`
	User       User       `json:"user"`
	Attributes Attributes `json:"object_attributes"`
}

func ExtractTagFromRef(ref string) string {
	tokens := strings.Split(ref, "/")
	return tokens[len(tokens)-1]
}
