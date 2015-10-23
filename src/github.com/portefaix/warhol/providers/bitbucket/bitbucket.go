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

package bitbucket

// User defines the Bitbucket user object
type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// Repository defines Bitbucket user object
type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"fill_name"`
	UUID     string `json:"uuid"`
}

// type Target struct {
// 	Type    string `json:"type,omitempty"`
// 	Hash    string `json:"hash,omitempty"`
// 	Author  string `json:"author,omitempty"`
// 	Message string `json:"message,omitempty"`
// 	Date    string `json:"date,omitempty"`
// }

// type New struct {
// 	Name   string `json:"name,omitempty"`
// 	Type   string `json:"type,omitempty"`
// 	Target Target `json:"target,omitempty"`
// }

// Commit defines the Bitbucket commit object
type Commit struct {
	Hash    string `json:"hash,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Author  string `json:"author,omitempty"`
}

// Changes defines the Bitbucket changes object
type Changes struct {
	// New New `json:"new,omitempty"`
	Created   bool     `json:"created"`
	Forced    bool     `json:"forced"`
	Close     bool     `json:"closed"`
	Truncated bool     `json:"truncated"`
	Commits   []Commit `json:"commits"`
}

// Push defines the Bitbucket push object
type Push struct {
	Changes []Changes `json:"changes"`
}

// PushWebHook defines the Bitbucket webhook object for push
type PushWebHook struct {
	Actor      User       `json:"actor"`
	Repository Repository `json:"repository"`
	Push       Push       `json:"push"`
}
