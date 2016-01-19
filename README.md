# minecraft-proxy

This is a HTTP to WebSockets proxy written in Go. It is not meant as a general purpose proxy, but rather as a proxy between the `minecraft-commander` and `mineslave` projects.

## Why?
The `minecraft-commander` code was migrated from using WebSockets to HTTP calls to communicate with the `mineslave` Bukkit plugin. The `mineslave` plugin would have to be updated to support HTTP calls, but the web server implementations in Java seemed too complicated, while Go's implementations seemed much cleaner. This project exists to allow `minecraft-commander` to use HTTP calls, while for `mineslave` to continue using the old WebSocket-based implementation.

## What it does
The code here operates as both a web server and as a WebSocket client.

When a HTTP POST request is received, a WebSocket connection is initiated to the `mineslave` plugin running on `localhost:8080`. The player name is transmitted, followed by the request's body. Afterwards, we wait for both transmissions' return messages and send back the second return message for the HTTP client (`minecraft-commander`).
