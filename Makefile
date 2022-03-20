.PHONY: test

export GOMAPSROOT=${PWD}

start:	
		docker-compose -f util/docker/docker-compose.yaml up

test:
		go test -v ./test
