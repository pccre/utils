package Mut

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WS struct {
	WS  *websocket.Conn
	Mut *sync.Mutex
}

func (c *WS) WriteJSON(content interface{}) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}
	return c.WriteRaw(data)
}

func (c *WS) WriteRaw(content []byte) error {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	return c.WS.WriteMessage(1, content)
}
