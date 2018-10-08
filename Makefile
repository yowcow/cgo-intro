BIN := main

all: $(BIN)
	ldd $(BIN) || true

$(BIN):
	GO111MODULE=on go build --ldflags '-extldflags "-static"' -x -o $@

test:
	GO111MODULE=on go test -v ./

clean:
	rm -rf $(BIN)

.PHONY: all test clean
