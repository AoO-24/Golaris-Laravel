BINARY_NAME=celeritasApp.exe

## build: builds all binaries
build:
    @go mod vendor
    @go build -o tmp/${BINARY_NAME} .
	@echo Celeritas built!

run:
	@echo Staring Celeritas...
    @start /min cmd /c tmp\${BINARY_NAME} &
	@echo Celeritas started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
	
stop:
	@echo "Starting the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Celeritas

restart: stop start