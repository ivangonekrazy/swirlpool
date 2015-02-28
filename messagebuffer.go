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
