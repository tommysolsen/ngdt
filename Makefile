.PHONY: clean build install

build:
	go generate
	go-bindata -o assets/templates_temp.go -pkg assets static/*
	go build -ldflags="-s -w" -o bin/nf *.go

run: build
	bin/nf labels show

clean:
	rm -f ./**/*_temp.go
	rm -rf bin/*

install: build
	cp bin/nf /usr/local/bin
	chmod u+x /usr/local/bin/nf
