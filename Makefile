bench:
	go test -bench . -benchmem -test.timeout 60m

test:
	go vet
	go test -v . -coverpkg ./...

