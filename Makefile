BINARY=conncurrencyApp

build:
	CGO_ENABLED=0 go build -o ${BINARY}  main.go

run: build
	./${BINARY}
