up: 
	cd app && go run main.go

test:
	go test -v ./... -count=1
