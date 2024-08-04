up: 
	cd app && go run main.go

test:
	go test -cover ./... -count=1

testv:
	go test -v ./... -count=1 -cover

build:
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-o engine \
		./app/
	@ echo "done"
