package main

import "runtime/trace"

type room struct {
	// forward is a channel that holds incoming messages
	// that should be for the other clients.
	forward chan []byte

	// join is a channel for clients wishing to join the room.
	join chan *client

	// leave is a channel for clients wishing to leave the room.
	leave chan *client

	// clients holds all current clients in this room.
	clients map[*client]bool

	// tracer will receive trace information of activity in the room.
	tracer trace.Tracer
}