LIB := lib
BUILD := build

all: $(LIB) $(LIB)/libmessage.so

$(LIB)/libmessage.so: $(BUILD)/libmessage.o
	$(CC) -shared -Wall -O2 -I ./include -fPIC $< -o $@

$(LIB):
	mkdir -p $@

$(BUILD)/%.o: $(BUILD)
	$(CC) -c -Wall -O2 -I ./include $*.c -o $@

$(BUILD):
	mkdir -p $@

clean:
	rm -rf $(BUILD) $(LIB)

test: $(BUILD)/libmessage-test
	./$<

.PHONY: all test clean
