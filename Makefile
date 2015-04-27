PID = tmp/server.pid
GO_FILES = $(wildcard server/*.go)

local: clean
	make restart
	fswatch -0 server/*.go | xargs -0 -n 1 -I {} make restart || make kill

kill:
	[ -f $(PID) ] && kill -9 `cat $(PID)` || true
	make clean

restart:
	make kill
	cd server ; go build
	./server/server -port 9393 & echo $$! > $(PID)

clean:
	rm -f server/server
	rm -f $(PID)

.PHONY: serve restart kill clean
