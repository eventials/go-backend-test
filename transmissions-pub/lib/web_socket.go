package lib

import "github.com/gorilla/websocket"

var WbUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
