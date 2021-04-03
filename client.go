package ddp

import (
	"github.com/gorilla/websocket"
)

type Config struct {
	// Dialer [optional]
	Dialer *websocket.Dialer
}
