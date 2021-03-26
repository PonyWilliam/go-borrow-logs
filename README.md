# BorrowLogs Service

This is the BorrowLogs service

Generated with

```
micro new --namespace=go.micro --type=service borrowLogs
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.borrowLogs
- Type: service
- Alias: borrowLogs

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./borrowLogs-service
```

Build a docker image
```
make docker
```