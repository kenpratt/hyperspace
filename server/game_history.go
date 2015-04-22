package main

import (
	"container/list"
)

type StateMutationFunction func(*GameState) error

type GameHistory struct {
	events *list.List
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

func (h *GameHistory) Exec(e GameEvent) error {
	return e.Execute(h.CurrentState())
}

func (h *GameHistory) Update() error {
	// get last game state
	state := h.CurrentState()

	// calculate time since last update (in milliseconds)
	now := MakeTimestamp()
	elapsed := now - state.Time

	// update game state
	h.events.PushBack(state.Tick(now, elapsed))
	return nil
}

func (h *GameHistory) CurrentState() *GameState {
	return h.events.Back().Value.(*GameState)
}
