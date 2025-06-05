INCLUDE=src/*
EXECUTABLE_OUTPUT=out/main
MAIN_SRC=src/main.go

build:
	GOARCH=amd64 GOOS=darwin go build -o ${EXECUTABLE_OUTPUT}-darwin ${MAIN_SRC}
	GOARCH=amd64 GOOS=linux go build -o ${EXECUTABLE_OUTPUT}-linux ${MAIN_SRC}
	GOARCH=amd64 GOOS=windows go build -o ${EXECUTABLE_OUTPUT}-windows ${MAIN_SRC}

run:
	go run ${MAIN_SRC}

clean:
	go clean
	rm ${EXECUTABLE_OUTPUT}-darwin
	rm ${EXECUTABLE_OUTPUT}-linux
	rm ${EXECUTABLE_OUTPUT}-windows

compile:
	go vet ${INCLUDE}
	go fmt ${INCLUDE}
	make build