PHONY: run-docker 
run-docker:
		docker-compose up -d --build


.PHONY: run
run:
		while read line; do export $line; done < env.local
		go run ./cmd/app/main.go

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: fmt 
fmt:
		go fmt ./... 

.DEFAULT_GOAL := run 

