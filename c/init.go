package c // constants

import (
	"github.com/gofiber/contrib/websocket"
	jsoniter "github.com/json-iterator/go"
)

var JSON = jsoniter.ConfigFastest
var WSConfig = websocket.Config{EnableCompression: true}
