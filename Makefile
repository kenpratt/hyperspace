all: local

.PHONY: local
local:
	cd public; python -m SimpleHTTPServer 9393
