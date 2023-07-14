.PHONY:

build:
	go build -o ./.bin/bot  cmd/bot/main.go

run:
	docker-compose up --build

lint:
	golangci-lint --config .golangci.yml run ./... --deadline=2m --timeout=2m

build-image:
	docker build -t quizdocker .

start-container:
	docker run --name QuizContainer -p 80:80 quizdocker

prune:
	docker container prune -f