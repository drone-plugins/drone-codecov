# plugin-codecov

Woodpecker plugin to send coverage reports to [Codecov](https://codecov.io/).

## Build

Build the binary with the following command:

```sh
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

go build -ldflags '-s -w -extldflags "-static"' -o plugin-codecov
```

## Docker

Build the Docker image with the following command:

```sh
docker build -f docker/Dockerfile.alpine -t woodpeckerci/plugin-codecov:next-alpine .
```
