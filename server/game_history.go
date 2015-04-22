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

func (h *GameHistory) Tick() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	// get last game state
	oldState := h.currentState()

	// update game state
	newState := oldState.Tick(MakeTimestamp())
	h.events.PushBack(newState)
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
