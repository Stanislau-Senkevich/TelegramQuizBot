.PHONY:

build:
	go build -o ./.bin/bot  cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t quizdocker .

start-container:
	docker run --name QuizContainer -p 80:80 quizdocker

prune:
	docker container prune -f