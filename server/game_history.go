package main

import (
	"container/list"
	"sync"
)

type StateMutationFunction func(*GameState) error

type GameHistory struct {
	events *list.List
	mu     sync.Mutex
}

type HistoryEntry struct {
	time uint64
	fn   StateMutationFunction
}

func CreateGameHistory() *GameHistory {
	events := list.New()
	events.PushBack(CreateGameState(MakeTimestamp()))
	return &GameHistory{
		events: events,
	}
}

func (h *GameHistory) Run(e GameEvent) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	return e.Execute(h.currentState())
}

func (h *GameHistory) Update() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	// get last game state
	state := h.currentState()

	// calculate time since last update (in milliseconds)
	now := MakeTimestamp()
	elapsed := now - state.Time

	// update game state
	h.events.PushBack(state.Tick(now, elapsed))
	return nil
}

func (h *GameHistory) CurrentState() *GameState {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.currentState()
}

func (h *GameHistory) currentState() *GameState {
	return h.events.Back().Value.(*GameState)
}
