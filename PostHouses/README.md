# PostHouses Service

This is the PostHouses service

Generated with

```
micro new renting/PostHouses --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.PostHouses
- Type: srv
- Alias: PostHouses

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./PostHouses-srv
```

Build a docker image
```
make docker
```