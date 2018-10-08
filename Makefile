BIN := main

all: $(BIN)
	ldd $(BIN) || true

$(BIN):
	go build --ldflags '-extldflags "-static"' -x -o $@

test:
	go test -v ./

clean:
	rm -rf $(BIN)

.PHONY: all test clean
