package socket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Hub interface {
	RegisterClient(id uint64, conn *websocket.Conn)
	Broadcast(senderID uint64, recipientIDs []uint64, message interface{})
	Run()
	GetClients() map[uint64]*client
}

type hub struct {
	clients         map[uint64]*client
	outboundPackets chan *packet
	register        chan *client
}

type packet struct {
	senderID     uint64
	recipientIDs []uint64
	message      interface{}
}

type client struct {
	id   uint64
	conn *websocket.Conn
}

var globalHub Hub

func InitHub() {
	if globalHub != nil {
		return
	}

	globalHub = &hub{
		clients:         make(map[uint64]*client),
		outboundPackets: make(chan *packet),
		register:        make(chan *client),
	}

	go globalHub.Run()
}

func GetHub() Hub {
	return globalHub
}

func (h *hub) RegisterClient(
	id uint64,
	conn *websocket.Conn,
) {
	c := &client{
		id:   id,
		conn: conn,
	}

	h.register <- c
}

func (h *hub) Broadcast(
	senderID uint64,
	recipientIDs []uint64,
	message interface{},
) {
	go func() {
		packet := &packet{
			senderID:     senderID,
			recipientIDs: recipientIDs,
			message:      message,
		}

		h.outboundPackets <- packet
	}()
}

func (h *hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.id] = client
		case message := <-h.outboundPackets:
			h.writeToClient(message)
		}
	}
}

func (h *hub) GetClients() map[uint64]*client {
	return h.clients
}

func (h *hub) writeToClient(packet *packet) {
	for _, id := range packet.recipientIDs {
		client, ok := h.clients[id]
		if !ok {
			continue
		}

		conn := client.conn

		bytes, _ := json.Marshal(packet.message)

		err := conn.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			conn.Close()
			delete(h.clients, id)
		}
	}
}
