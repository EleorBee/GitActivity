run: build
	./bin/GitActivity.exe

build:
	go build -o bin/GitActivity.exe GitActivity/cmd/github-activity

install:
	go install GitActivity/cmd/github-activity