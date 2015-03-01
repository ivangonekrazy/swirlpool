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
			h.connections[c] = true
			fmt.Printf("[ + ] Connection: %v (%v)\n", c, len(h.connections))
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.messageChan)
			}
			fmt.Printf("[ - ] Connection: %v (%v)\n", c, len(h.connections))
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.messageChan <- m:
					fmt.Printf("Delivered %v to %v\n", m, c)
				default:
					fmt.Printf("Stuck connection? %v\n", c)
					delete(h.connections, c)
					close(c.messageChan)
					continue
				}
			}

		}
	}
}
