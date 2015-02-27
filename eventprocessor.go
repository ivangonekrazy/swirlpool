package main

import (
	"fmt"
)

type EventsProcessor struct {
	messages MessageBuffer
	intake   chan Message
}

func (p *EventsProcessor) Run() {
	for {
		select {
		case m := <-p.intake:
			p.messages.PushBack(m)
			fmt.Printf("%v - %v\n", p.messages.Len(), m.String())

			// if any data payload matches my name, emit a new broadcast
			for _, v := range m.data {
				if v == "ivan" {
					h.broadcast <- NewMessage("Welcome, Ivan!")
				}
			}
		}
	}
}
