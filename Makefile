up: 
	cd app && go run main.go

test:
	go test ./... -count=1 -coverprofile=c.out -covermode=atomic

testv:
	go test -v ./... -count=1 -coverprofile=c.out -covermode=atomic
