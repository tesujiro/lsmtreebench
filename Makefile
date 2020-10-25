bench:
	go test -bench . -benchmem -test.timeout 10m | tee ./result.txt

test:
	go vet
	go test -v . -coverpkg ./...

