# dockerhub-webhook

A simple server to add obuscated URLs, auth tokens, parse out and forward docker deploy requests to jenkins
with a namespace deployment request.

Written in [Golang.](http://golang.org)

## About

Webhooks from dockerhub are public events. We needed an alternative to Smee.io client to manage
these requests, add additional meta-data, and provide a token mechanism for additional security.

## Command Line Usage

```
Small server that obuscates a URL path and relays requests from dockerhub to jenkins

Usage:
  dockerhub-webhook [command]

Available Commands:
  help        Help about any command
  start       Starts the dockerhub webhook server
  version     Version of the application

Flags:
  -a, --alive-path string     Path to handle health checks (default "/v1.0/alive")
      --config string         config file (default is $HOME/.dockerhub-webhook)
  -D, --debug                 Debug
  -h, --help                  help for dockerhub-webhook
  -O, --hostname string       Host or IP for this server (default "0.0.0.0")
  -C, --max-conn int          Maximum conn for this server
  -X, --max-procs int         Maximum processors for this server
  -n, --namespace string      Namespace for pricessing requests (default "development")
  -y, --notify-path string    Path to handle notification events (default "/v1.0/notify")
  -L, --port int              Listen port for this server (default 8080)
  -P, --profile-port int      Profile port for this server (default 6060)
  -s, --status-path string    Path to get server status (default "/v1.0/status")
  -e, --target-host string    Host to relay request (default "jenkins")
  -g, --target-path string    Path to webhook (default "/generic-webhook-trigger/invoke/")
  -o, --target-port int       Port to access (default 8080)
  -t, --target-token string   Authentication token
  -v, --valid-tokens string   List of valid tokens to access notifications (comma delim)

Use "dockerhub-webhook [command] --help" for more information about a command.


```
## Configuration

A config file is mandatory. You can place it in the same directory as the
application or in your home directory. The name, with period, is:

.dockerhub-webhook.yaml

This file can be empty, but must exist.

You have two places you can configure the server:

* Pass parameters in your application
* Put attributes in your config file

Parameters or attributes can be values or can be paths to a file that contains
the value. These are named the same.

an example config file is included under /examples

## Building

This code was built with 1.12.6, but you can test with backwards versions.

. build.sh is the tool to create multiple executables. Edit what you need/don't need.

For package management, look to dep for instructions: <https://github.com/golang/dep>

commands:
```
dep init
dep ensure -add
dep ensure -update
dep status
```

Information on Golang installation, including pre-built binaries, is available at <http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `dockerhub-webhook` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

## Docker Image

The docker file and image provides a means to run this as a container under docker.

CLI Example:
```
# Build image
docker build --force-rm --no-cache --build-arg release_tag=1.0.0 -t composer22/dockerhub-webhook:latest .

# Run the service
docker service create --name dockerhub-webhook --replicas 1 \
-e "DW_VALID_TOKENS=tokenvalue1,tokenvalue2" \
-e "DW_TARGET_HOST=jenkins.mysite.com" \
-e "DW_TARGET_PORT=8080" \
-e "DW_TARGET_TOKEN=ABCDEFG" \
-p 8080:8080 \
composer22/dockerhub-webhook:1.0.0

# Where dockerhub-req.json is a file that contains example json from dockerhub webhook
curl -v -d @dockerhub-req.json -H "Content-Type: application/json" http://dockerhub-webhook/v1.0/notify?token=tokenvalue2


# Check the logs

```

## License

(The MIT License)

Copyright (c) 2019 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
