package main

// Ordered and bounded list of Messages
type MessageBuffer struct {
	messages []Message
	capacity int
}

func NewMessageBuffer(capacity int) MessageBuffer {
	mb := MessageBuffer{}
	mb.capacity = capacity
	return mb
}

func (mb *MessageBuffer) Reset() {
	mb.messages = nil
}

// Add a Message at the end of the list.
func (mb *MessageBuffer) PushBack(m Message) {
	// if the buffer is at capacity, remove the head element
	if mb.Len() == mb.capacity {
		mb.messages = mb.messages[1:]
	}

	mb.messages = append(mb.messages, m)
}

func (mb *MessageBuffer) Len() int {
	return len(mb.messages)
}

func (mb *MessageBuffer) List() []Message {
	return mb.messages
}

type Messagelength struct {
	message Message
	length  int
}

// Return the longest length at the end of the MessageBuffer.
func LatestStreak(mb *MessageBuffer) Messagelength {
	streak_len := 1
	for i := mb.Len() - 1; i > 0; i-- {
		if mb.messages[i].Same(mb.messages[i-1]) {
			streak_len++
		} else {
			return Messagelength{message: mb.messages[i], length: streak_len}
		}
	}

	return Messagelength{message: mb.messages[mb.Len()-1:][0], length: streak_len}
}
