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

package main

import (
	"flag"
	"fmt"
	"log"

	//log "github.com/Sirupsen/logrus"

	"github.com/portefaix/warhol/api"
	"github.com/portefaix/warhol/logging"
	"github.com/portefaix/warhol/providers/docker"
)

var (
	port            string
	debug           bool
	version         bool
	dockerHost      string
	dockerTLSVerify bool
	dockerCertPath  string
	registryURL     string
	username        string
	password        string
	email           string
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "port to use")
	flag.StringVar(&dockerHost, "docker-host", "unix:///var/run/docker.sock", "address of Docker host")
	flag.BoolVar(&dockerTLSVerify, "docker-tls-verify", false, "use TLS client for Docker")
	flag.StringVar(&dockerCertPath, "docker-cert-path", "", "path to the cert.pem, key.pem, and ca.pem for authenticating to Docker")
	flag.StringVar(&registryURL, "registry-url", "127.0.0.1:5000", "host:port of the registry for pushing images")
	flag.StringVar(&username, "username", "", "Username used for Docker registry")
	flag.StringVar(&password, "password", "", "Password used for Docker registry")
	flag.StringVar(&password, "email", "", "Email used for Docker registry")
	flag.Parse()
}

func getDockerBuilder() (*docker.Builder, error) {
	return docker.NewBuilder(
		dockerHost,
		dockerTLSVerify,
		dockerCertPath,
		registryURL,
		&docker.Authentication{
			Username: username,
			Password: password,
			Email:    email,
		})
}

func main() {
	if debug {
		//log.SetLevel(log.DebugLevel)
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	if version {
		fmt.Printf("Warhol v%s\n", Version)
		return
	}
	log.Print("[INFO] [warhol] Creates the Docker builder")
	builder, err := getDockerBuilder()
	if err != nil {
		log.Printf("[FATAL] [warhol] Error with Docker : %v", err)
		return
	}
	go builder.Build()
	go builder.Push()
	e := api.GetWebService(builder)
	if debug {
		e.Debug()
		builder.Debug()
	}

	log.Printf("[INFO] [warhol] Warhol is ready on %s", port)
	e.Run(fmt.Sprintf(":%s", port))
}
