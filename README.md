# Service API

A quick'n dirty HTTP API for starting & stopping linux services.


## Setup

Build serviceapi:

    go build

Edit config file:

  - HTTPPort: port to listen for HTTP requests
  - AllowedServices: list of remote-controlable services

## Running

    sudo ./serviceapi config.json

## Using

    curl http://<host ip>/<service name>/start
    curl http://<host ip>/<service name>/stop

Example:
  
    curl http://192.168.0.1/postgresql/start



## Security notice

This soft is absolutely not secured in any way, so make sure:

  - You only expose non-sensitive linux services
  - The API won't be accessible from the internet

