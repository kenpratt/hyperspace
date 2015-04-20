PID = tmp/server.pid
GO_FILES = $(wildcard server/*.go)

local: clean
	make restart
	fswatch -o . | xargs -n1 -I{} make restart || make kill

kill:
	kill `cat $(PID)` || true

restart:
	make kill
	cd server ; go build
	./server/server -debug -port 9393 & echo $$! > $(PID)

clean:
	rm -f server/server
	echo '' > tmp/server.pid

.PHONY: serve restart kill clean
