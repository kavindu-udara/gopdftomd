build:
	go build -o bin/pdf2md

run: build
	./bin/pdf2md
