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

	log "github.com/Sirupsen/logrus"

	"github.com/portefaix/warhol/api"
	"github.com/portefaix/warhol/providers/docker"
)

var (
	port    string
	debug   bool
	version bool
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "port to use")
	flag.Parse()
}

func main() {
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	if version {
		fmt.Printf("Warhol v%s\n", Version)
		return
	}
	e := api.GetWebService()
	if debug {
		e.Debug()
	}
	log.Info("Launch Docker builder")
	dockerBuilder, err := docker.NewDockerBuilder()
	if err != nil {
		log.Fatalf("Error with Docker : %v", err)
		return
	}
	log.Info("Build image")
	err = dockerBuilder.BuildImage("foo",
		"/home/nlamirault/Perso/Portefaix/warhol/Dockerfile")
	log.Infof("Docker err : %v", err)
	log.Infof("Launch Warhol on %s", port)
	e.Run(fmt.Sprintf(":%s", port))
}
