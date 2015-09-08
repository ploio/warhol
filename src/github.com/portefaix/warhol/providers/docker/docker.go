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

package docker

import (
	"bytes"
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

// DockerBuilder builds Docker images
type DockerBuilder struct {
	*docker.Client
}

// SOCKET represents the Docker socket endpoint
const SOCKET = "unix:///var/run/docker.sock"

// NewDockerBuilder creates a new instance of DockerBuilder
func NewDockerBuilder() (*DockerBuilder, error) {
	client, err := docker.NewClient(SOCKET)
	if err != nil {
		return nil, err
	}
	return &DockerBuilder{
		Client: client,
	}, nil
}

func getImageName(name string) string {
	return fmt.Sprintf("warhol/%s", name)
}

func (db *DockerBuilder) BuildImage(name string, dockerfile string) error {
	var buf bytes.Buffer
	opts := docker.BuildImageOptions{
		Name:           getImageName(name),
		Dockerfile:     "Dockerfile",
		Remote:         "github.com/nlamirault/aneto",
		SuppressOutput: true,
		OutputStream:   &buf,
	}
	return db.Client.BuildImage(opts)
}
