# hyperspace

A real-time multiplayer space shooter.

## Development

 * Run `make` to run a local copy at http://localhost:9393.

 * Run `./watch` to run a local copy that auto-restarts when server files are changed (requires fswatch - `brew install fswatch`)

## Installation

### Add to Nginx configuration

Add to bottom of http block:

```conf
include /home/hyperspace/hyperspace/etc/nginx.conf;
```

Add systemd service:

```sh
sudo ln -s /home/hyperspace/hyperspace/etc/hyperspace.service /etc/systemd/system/
sudo systemctl start hyperspace
```

### Add ability to restart server:

```sh
sudo EDITOR=emacs visudo
```

```
hyperspace ALL=(ALL) NOPASSWD: /home/hyperspace/hyperspace/bin/restart
```

### Add dependencies

```sh
go get github.com/gorilla/websocket
```
