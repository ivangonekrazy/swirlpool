package main

import (
	"fmt"
)

type EventsProcessor struct {
	messages MessageBuffer
	intake   chan Message
}

func (p *EventsProcessor) Run(broadcastChan chan Message) {
	for {
		select {
		case m := <-p.intake:
			p.messages.PushBack(m)
			broadcastChan <- m
			fmt.Printf("Processing %v - %v\n", m.String(), p.messages.Len())

			// detect and announce the kill streaks from Unreal Tournament
			s := LatestStreak(&p.messages)
			fmt.Printf("----------- %v\n", s)
			switch s.length {
			case 8:
				broadcastChan <- NewMessage("HOLY SHIT!", "killingspree")
			case 7:
				broadcastChan <- NewMessage("Ludicrous Kill", "killingspree")
			case 6:
				broadcastChan <- NewMessage("Monster Kill", "killingspree")
			case 5:
				broadcastChan <- NewMessage("Ultra Kill", "killingspree")
			case 4:
				broadcastChan <- NewMessage("Mega Kill", "killingspree")
			case 3:
				broadcastChan <- NewMessage("Multi Kill", "killingspree")
			case 2:
				broadcastChan <- NewMessage("Double Kill", "killingspree")
			}
		}
	}
}
