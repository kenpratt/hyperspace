all: local

.PHONY: local
local:
	cd server/ ; go build
	./server/server -debug

.PHONY: clean
clean:
	rm -f server/server
