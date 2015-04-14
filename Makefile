all: local

.PHONY: local
local:
	go run server/*.go
