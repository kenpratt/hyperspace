# hyperspace

A real-time multiplayer space shooter

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
