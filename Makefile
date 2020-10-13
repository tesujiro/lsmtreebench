bench:
	go test -bench . -benchmem -test.timeout 10m

test:
	go vet
	go test -v . -coverpkg ./...

