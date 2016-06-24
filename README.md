# flowd

Flowd is a simple network monitoring service. It captures TCP/IP traffic and keeps track of
`source -> destination` mappings, which are exposed via HTTP server running on port `7777`.

Sample output:

```json
{
  "10.0.2.15:37745": "128.101.240.215:80",
  "10.0.2.15:39323": "173.220.95.85:80",
  "10.0.2.15:44995": "5.153.231.35:80",
  "10.0.2.15:50510": "128.61.240.89:80",
  "10.0.2.15:58835": "64.50.233.100:80",
  "128.101.240.215:80": "10.0.2.15:37745",
  "128.61.240.89:80": "10.0.2.15:50510",
  "173.220.95.85:80": "10.0.2.15:39323",
  "5.153.231.35:80": "10.0.2.15:44995",
  "64.50.233.100:80": "10.0.2.15:58835"
}
```

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
