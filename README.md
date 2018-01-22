# drone-codecov

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-codecov/status.svg)](http://beta.drone.io/drone-plugins/drone-codecov)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-codecov?status.svg)](http://godoc.org/github.com/drone-plugins/drone-codecov)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-codecov)](https://goreportcard.com/report/github.com/drone-plugins/drone-codecov)
[![](https://images.microbadger.com/badges/image/plugins/codecov.svg)](https://microbadger.com/images/plugins/codecov "Get your own image badge on microbadger.com")

Drone plugin to send coverage reports to [Codecov](https://codecov.io/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-codecov/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-codecov
docker build --rm -t plugins/codecov .
```

### Usage

```
docker run --rm \
  -e PLUGIN_TOKEN=xxx \
  -e DRONE_COMMIT=5f17090 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/codecov
```
