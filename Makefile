.PHONY: test

test:
	go test ./...

bench:
	go test -run=XXX -bench=. -benchmem ./...
