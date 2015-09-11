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
	"errors"
	"fmt"
	"io"

	// log "github.com/Sirupsen/logrus"
	"github.com/fsouza/go-dockerclient"
)

const (
	// SOCKET represents the Docker socket endpoint
	SOCKET = "unix:///var/run/docker.sock"

	// REGISTRY is the default Docker registry
	REGISTRY = "index.docker.io"
)

var (
	// ErrDockerAuthentication Can't authenticate to Docker
	ErrDockerAuthentication = errors.New("Docker authentication failed")
)

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

// Authentication represents authentication option for Docker registry
type Authentication struct {
	Username string
	Password string
	Email    string
}

// NewBuilder creates a new instance of DockerBuilder
func NewBuilder(host string, tls bool, certPath string, registryURL string, auth *Authentication) (*Builder, error) {
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
	authConf := docker.AuthConfiguration{
		Username:      auth.Username,
		Password:      auth.Password,
		Email:         auth.Email,
		ServerAddress: registryURL,
	}
	log.Printf("[INFO] [docker] Check Docker authentication: %v", authConf)
	if client.AuthCheck(&authConf) != nil {
		return nil, ErrDockerAuthentication
	}
	builder := &Builder{
		Client:      client,
		RegistryURL: registryURL,
		AuthConfig:  authConf,
		BuildChan:   make(chan *Project),
		PushChan:    make(chan *Project),
	}
	log.Printf("[DEBUG] [docker] Creating Docker builder : %#v", builder)
	env, err := builder.Client.Version()
	if err != nil {
		log.Printf("[WARN] [docker] Can't retrieve Docker version: %v",
			err)
	} else {
		log.Printf("[INFO] [docker] %v", env)
	}

	return builder, nil
}

// Project represents a Git project
type Project struct {
	Name   string
	Remote string
}

// NewProject creates a new instance of Project
func (db *Builder) NewProject(name string, dockerfile string, remote string) *Project {
	return &Project{Name: name, Remote: remote}
}

// GetImageName returns the Docker image name (depends on Registry URL)
func (db *Builder) GetImageName(name string) string {
	if db.RegistryURL != REGISTRY {
		return fmt.Sprintf("%s/warhol/%s", db.RegistryURL, name)
	}
	return fmt.Sprintf("warhol/%s", name)
}

// Debug check Docker informations
func (db *Builder) Debug() {
	env, err := db.Client.Info()
	if err != nil {
		log.Printf("[WARN] [docker] Can't retrieve Docker informations: %v",
			err)
		return
	}
	log.Printf("[DEBUG] [docker] %#v", env)
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
	imageName := db.GetImageName(project.Name)
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
		Remote:       project.Remote,
		OutputStream: outputbuf,
	}
	log.Printf("[DEBUG] [docker] Building image : %s", imageName)
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
	imageName := db.GetImageName(project.Name)
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
	log.Printf("[DEBUG] [docker] Pushing image : %s", imageName)
	err := db.Client.PushImage(opts, db.AuthConfig)
	if err != nil {
		log.Printf("[ERROR] [docker] Can't push image %s : %v", imageName, err)
		return err
	}
	log.Printf("[INFO] [docker] Push image done : %s", imageName)
	return nil
}
