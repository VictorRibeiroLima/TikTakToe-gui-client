package internal

import (
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var addr = "tik-tak-toe-363515.rj.r.appspot.com"

func Connect(path string) (*websocket.Conn, *http.Response, error) {
	var u url.URL
	u = url.URL{Scheme: "ws", Host: addr, Path: path}

	return websocket.DefaultDialer.Dial(u.String(), nil)

}
