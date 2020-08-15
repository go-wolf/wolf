package main

import (
	"fmt"
	"wolf/network/ws"
)

func main() {
	wsServer := ws.NewWsServer("127.0.0.1", 8000)
	wsServer.SetOnOpen(OnOpen)
	wsServer.SetOnMessage(OnMessage)
	wsServer.SetOnClose(OnClose)
	wsServer.Start()
}
func OnOpen(wsConn *ws.WsConn) {
	fmt.Println(1)
}
func OnMessage(wsConn *ws.WsConn, msg []byte) {
	fmt.Println(2)
}
func OnClose(wsConn *ws.WsConn) {
	fmt.Println(3)
}
