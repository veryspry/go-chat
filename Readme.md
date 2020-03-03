### Websocket messenger server written in Go

This is a room based messenger server with message persistance.

All websocket connections are handled by one endpoint `ws/{roomID}`. There are three main vehicles to handle "room" separation: `Hub`, `Room`, and `Client`. `Hub` is the top level room mananger. It holds all of the open rooms and manages room creation and deletion. `Room` holds information about each reoom inculding a list of `Clients`. `Clients` are one "entity" in a room. This includes their `*websocket.Conn` and their uniqueID.

Authentication for HTTPS connections is handled by JWT bearer token headers. And the websocket authentication will be handled by a ticketing system.


### Postgres Instsall

Installing postgres on Ubuntu:
https://linuxize.com/post/how-to-install-postgresql-on-ubuntu-18-04/

- user
- pw
- db name
- host

start database: `pg_ctlcluster 11 main start`
