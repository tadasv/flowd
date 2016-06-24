dev:
	docker run --net=host -p 7777:7777 --rm -ti -v $(PWD)/src:/go/src/github.com/tadasv/flowd golang:1.6.2 bash

docker-build:
	docker run --rm -w /go/src/github.com/tadasv/flowd -v $(PWD)/src:/go/src/github.com/tadasv/flowd golang:1.6.2 bash -c "apt-get update -y && apt-get install -y libpcap-dev && go get -v && go build --ldflags '-extldflags \"-static\"'"

docker-image:
	docker build -t tadasv/flowd .
