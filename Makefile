build: deps
	go build
deps:
	glide install

install: deps
	go install
