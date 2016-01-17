deps:
	@go get ./...

copy: deps
	@cp tomato.png ${GOBIN}/tomato.png

build: deps
	@go install ./...

lint: build
	@go vet ./...
	golint ./...

test: build
	@go test ./...

pack: copy build;
