all: embed compile

embed:
	./embed.sh

compile:
	go get
	go build