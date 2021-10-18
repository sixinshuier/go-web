export tag=v1.0
root:
	export ROOT=github.com/sixinshuier/go-web

build:
	echo "building http server binary"
	mkdir -p output/bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/bin .

release: build
	echo "building httpserver container"
	docker build -t sixinshuier/go-web:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push sixinshuier/go-web:v1.0