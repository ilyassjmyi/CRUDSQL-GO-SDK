package openapi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

// WebSocketClient represents a WebSocket client
type WebSocketClient struct {
	conn *websocket.Conn
}

// NewWebSocketClient creates a new WebSocket client
func NewWebSocketClient(ctx context.Context, baseURL, key, model, event string) (*WebSocketClient, error) {
	u := url.URL{Scheme: "ws", Host: baseURL, Path: "/api/ws/" + model + "/" + event}
	headers := http.Header{}
	headers.Add("Authorization", key)
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return &WebSocketClient{conn: conn}, nil
}

// Listen listens for incoming WebSocket messages
func (c *WebSocketClient) Listen(handler func(event, model string, data interface{})) {
	defer c.conn.Close()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("unmarshal:", err)
			continue
		}

		event, _ := msg["event"].(string)
		model, _ := msg["model"].(string)
		data := msg["data"]

		handler(event, model, data)
	}
}
