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
	"bufio"
	"log"
	//"bytes"
	"fmt"
	"io"

	// log "github.com/Sirupsen/logrus"
	"github.com/fsouza/go-dockerclient"
)

// SOCKET represents the Docker socket endpoint
const SOCKET = "unix:///var/run/docker.sock"

// Builder builds Docker images
type Builder struct {

	// Docker client instance
	*docker.Client

	// Registry server to push the image
	RegistryURL string

	// Represents the authentication in the Docker index server
	AuthConfig docker.AuthConfiguration

	// Channel to build image
	BuildChan chan *Project

	// Channel to push image
	PushChan chan *Project
}

// NewBuilder creates a new instance of DockerBuilder
func NewBuilder(host string, tls bool, certPath string, registryURL string) (*Builder, error) {
	var client *docker.Client
	var err error
	if tls {
		cert := fmt.Sprintf("%s/cert.pem", certPath)
		key := fmt.Sprintf("%s/key.pem", certPath)
		ca := fmt.Sprintf("%s/ca.pem", certPath)
		client, err = docker.NewTLSClient(host, cert, key, ca)
	} else {
		client, err = docker.NewClient(host)
	}
	if err != nil {
		return nil, err
	}
	return &Builder{
		Client:      client,
		RegistryURL: registryURL,
		AuthConfig:  docker.AuthConfiguration{},
		BuildChan:   make(chan *Project),
		PushChan:    make(chan *Project),
	}, nil
}

// Project represents a Git project
type Project struct {
	Name       string
	Dockerfile string
	Remote     string
}

// NewProject creates a new instance of Project
func (db *Builder) NewProject(name string, dockerfile string, remote string) *Project {
	return &Project{Name: name, Dockerfile: "Dockerfile", Remote: remote}
}

func getImageName(name string) string {
	return fmt.Sprintf("warhol/%s", name)
}

// ToPipeline send a project to build pipeline
func (db *Builder) ToPipeline(project *Project) error {
	log.Printf("[INFO] [docker] Send project to pipeline : %v", project)
	db.BuildChan <- project
	return nil
}

// Build read channel and build Docker image
func (db *Builder) Build() error {
	project := <-db.BuildChan
	log.Printf("[INFO] [docker] Start building project : %v", project)
	imageName := getImageName(project.Name)
	logsReader, outputbuf := io.Pipe()
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			log.Printf("[DEBUG] [docker] %s", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Printf("[WARN] [docker] There was an error with the scanner in attached container %v", err)
		}
	}(logsReader)

	opts := docker.BuildImageOptions{
		Name:         imageName,
		Dockerfile:   project.Dockerfile,
		Remote:       "github.com/nlamirault/aneto", //project.Remote,
		OutputStream: outputbuf,
	}

	err := db.Client.BuildImage(opts)
	if err != nil {
		log.Printf("[ERROR] [docker] Can't build image %s : %v", imageName, err)
		return err
	}
	log.Printf("[INFO] [docker] Build image done : %s", imageName)
	db.PushChan <- project
	return nil
}

// Push read channel and push to registry the new image
func (db *Builder) Push() error {
	project := <-db.PushChan
	log.Printf("[INFO] [docker] Start pushing project : %v", project)
	imageName := getImageName(project.Name)
	logsReader, outputbuf := io.Pipe()
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			log.Printf("[DEBUG] [docker] %s", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Printf("[WARN] [docker] There was an error with the scanner in attached container %v", err)
		}
	}(logsReader)

	opts := docker.PushImageOptions{
		Name:         imageName,
		Tag:          "latest",
		Registry:     db.RegistryURL,
		OutputStream: outputbuf,
	}

	err := db.Client.PushImage(opts, db.AuthConfig)
	if err != nil {
		log.Printf("[ERROR] [docker] Can't push image %s : %v", imageName, err)
		return err
	}
	return nil
}
