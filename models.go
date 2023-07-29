package main

type Server struct {
	State GameState
}

type EventType int64

const (
	Login EventType = 0
)

type Event struct {
	Resource string
	Data     []byte
}

type Player struct {
	Money int
}

type GameState struct {
	Players map[string]Player
}
