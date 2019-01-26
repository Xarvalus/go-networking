# go-networking
Research & Experiments with Go for networking & data with use of concurrency.

## Socket TCP Server
`server/`

Networking connections between Server & Client via TCP/IP Socket layer.

### RPC
`server/rpc/` 

Remote Procedure Call on file reads with use of `net.rpc` on TCP (and extra concurrent HTTP service).

### TCP Socket
`server/socket/`

Simple Chat between Server <-> Client via TCP connection.

## HTTP WebServer
`webserver/`

Basic MVC REST Web Server with use of GORM ORM & Postgres SQL Database. 

### HTTP REST API
`webserver/controllers/api`

REST API endpoints, fetching from Database & inserting with GORM.

### HTTP WebSockets
`webserver/controllers/websockets`

WebSocket connection, Echo service & Database fetching response.
