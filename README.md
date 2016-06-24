# flowd

Flowd is a simple network monitoring service. It captures TCP/IP traffic and keeps track of
`source -> destination` mappings, which are exposed via HTTP server running on port `7777`.

## Running on Docker

`docker run -p 7777:7777 --rm --net=host tadasv/flowd`

## Running everywhere else

`./flowd`

## Developing

To make development easier it comes docker dev shell.

Run `make dev` to start dev environment inside Docker.

Alternatively you will need libpcap-dev and golang to build the code.

## Building

Run `make docker-build` to build `flowd` from source code ussing Docker golang image.

## Creating docker image

- `make docker-build`
- `make docker-image`
