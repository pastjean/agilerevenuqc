server-linux: main.go
	GOOS=linux go build -o server-linux .

.PHONY: docker
docker: server-linux
	docker build -t pastjean/agileqc:v1 .
	docker push pastjean/agileqc:v1