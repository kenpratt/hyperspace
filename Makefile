all: local

.PHONY: local
local:
	go run server/main.go
