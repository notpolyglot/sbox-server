package main

type EventType int64

const (
	Login EventType = 0
)

type Event struct {
	Type EventType
}

type Player struct {
	Money int
}

type State struct {
	Players map[string]Player
}
