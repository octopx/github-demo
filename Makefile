.PHONY: clean
clean:
	rm -f *.out

test:
	go test -v ./...

coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...

benchmark:
	go test -bench=.
