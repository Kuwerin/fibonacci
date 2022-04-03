run:
		docker-compose up -d --build

check:
	golangci-lint run -c .golangci.yml --fix

.DEFAULT_GOAL := run

