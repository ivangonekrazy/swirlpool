package main

import (
	"fmt"
)

// message fan-out broker
type Hub struct {
	connections map[*Connection]bool // Registered connections.
	broadcast   chan Message         // Inbound messages from the connections.
	register    chan *Connection     // Register requests from the connections.
	unregister  chan *Connection     // Unregister requests from connections.
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			fmt.Printf("Register connection: %s (%s)\n", c, len(h.connections))
			h.connections[c] = true
		case c := <-h.unregister:
			fmt.Printf("Unregister connection: %s (%s)\n", c, len(h.connections))
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.messageChan)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.messageChan <- m:
				default:
					delete(h.connections, c)
					close(c.messageChan)
				}
			}
		}
	}
}
