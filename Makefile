build:
	docker build -t go-demo:latest -- .
run:
	docker run --rm -it -p "8181:8181" go-demo:latest
manual-build:
	docker run -v $(PWD):/go/src/go-demo -w /go/src/go-demo --rm -it golang:1.8 bash
