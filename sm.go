package nanosm

import (
	"context"
)

// State represents a state in the state machine.
type State string

// Event represents an event that triggers transitions.
type Event string

// StateMachine manages states, transitions, and actions.
type StateMachine struct {
	currentState State
	states       map[State]*StateConfig
	transitions  map[Event]*TransitionConfig
	// Define to retrieve state data from external storage
	stateFetcher func(ctx context.Context, arg ...any) State
}

// NewStateMachine initializes the state machine with the initial state.
func NewStateMachine(initialState State) *StateMachine {
	return &StateMachine{
		currentState: initialState,
		states:       make(map[State]*StateConfig),
		transitions:  make(map[Event]*TransitionConfig),
	}
}

// NewStateMachine initializes the state machine with the initial state.
func NewStateMachineWithExternalStorage(initialState State, stateFetcher func(ctx context.Context, args ...any) State) *StateMachine {
	return &StateMachine{
		currentState: initialState,
		states:       make(map[State]*StateConfig),
		transitions:  make(map[Event]*TransitionConfig),
		stateFetcher: stateFetcher,
	}
}

// State creates a new state in the state machine.
func (sm *StateMachine) State(state State) *StateConfig {
	if sm.states[state] == nil {
		sm.states[state] = &StateConfig{}
	}
	return sm.states[state]
}

// Transition creates a transition configuration for an event between two states.
func (sm *StateMachine) Transition(from State, to State) *TransitionConfig {
	return &TransitionConfig{
		fromState: from,
		toState:   to,
	}
}

// AddTransition registers a transition in the state machine.
func (sm *StateMachine) AddTransition(cfg *TransitionConfig) {
	sm.transitions[cfg.event] = cfg
}

func (sm *StateMachine) CurrentState(ctx context.Context) State {
	return sm.stateFetcher(ctx)
}

func (sm *StateMachine) fire(ctx context.Context, event Event, args ...interface{}) error {
	// Get current state of state machine
	currentState := sm.stateFetcher(ctx, args...)

	// // Check if a valid transition exists for the event.
	// transition, ok := sm.transitions[event]
	// if !ok || transition.fromState != sm.currentState {
	// 	return errors.New("invalid transition")
	// }

	// // Run the guard check for the transition, if defined.
	// if transition.guard != nil && !transition.guard(context.Background()) {
	// 	return errors.New("guard condition failed, transition not allowed")
	// }

	// // Exit action of the transition.
	// if transition.exitAction != nil {
	// 	transition.exitAction(ctx, args)
	// }

	// // Transition to the new state.
	// fmt.Printf("Transitioning from %s to %s on event %s\n", sm.currentState, transition.toState, event)
	// sm.currentState = transition.toState

	// // Entry action of the transition.
	// if transition.entryAction != nil {
	// 	transition.entryAction(ctx, args)
	// }
	// Exit action of the current state.
	currentConfig := sm.states[currentState]
	if currentConfig != nil && currentConfig.exitAction != nil {
		currentConfig.exitAction(ctx, args)
	}

	// Entry action of the new state.
	newConfig := sm.states[sm.currentState]
	if newConfig != nil && newConfig.entryAction != nil {
		newConfig.entryAction(ctx, args)
	}

	// Enter action of the new state.
	if newConfig != nil && newConfig.enterAction != nil {
		newConfig.enterAction(ctx, args)
	}

	return nil
}

// TriggerEvent processes an event and performs the appropriate state transition.
func (sm *StateMachine) TriggerEventCtx(ctx context.Context, event Event, args ...interface{}) error {
	return sm.fire(ctx, event, args...)
}

// TriggerEvent processes an event and performs the appropriate state transition.
func (sm *StateMachine) TriggerEvent(event Event, args ...interface{}) error {
	return sm.fire(context.Background(), event, args...)
}
