default: run

run:
	go run routes.go opsel.go

dist:
	rm -rf opsel
	go build -o opsel -ldflags "-s -w" -trimpath routes.go opsel.go