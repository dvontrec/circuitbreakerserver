all: runServer runClient
runServer: 
	@echo "running server"
	@go run cmd/restapi/main.go &
runClient:
	@echo "Starting client"
	@go run integrations/client/main.go
