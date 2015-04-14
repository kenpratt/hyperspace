all: local

.PHONY: local
local:
	go run server/*.go

.PHONY: clean
clean:
	rm -f server/server
