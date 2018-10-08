package main

/*
#include "./clang/libmessage.c"
*/
import "C"

type Message struct {
	ptr *C.char
}

func NewMessage(name string) *Message {
	input := (*C.char)(C.CString(name))
	ptr := C.alloc_message(input)
	return &Message{ptr}
}

func (m *Message) String() string {
	return C.GoString(m.ptr)
}

func (m *Message) Destroy() {
	C.free_message(m.ptr)
}
