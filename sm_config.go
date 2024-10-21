package nanosm

import "context"

type SmConfig struct {
	sm *StateMachine
}

// TransitionConfig holds transition details between states.
type TransitionConfig struct {
	fromState   State
	toState     State
	event       Event
	entryAction func(ctx context.Context, args ...any)
	exitAction  func(ctx context.Context, args ...any)
	enterAction func(ctx context.Context, args ...any)
	guard       func(ctx context.Context, args ...any) bool
}

// StateConfig holds the configuration for actions and guards on a state.
type StateConfig struct {
	entryAction ActionFunc
	exitAction  ActionFunc
	enterAction ActionFunc
	guard       func(ctx context.Context) bool
}

// Event sets the event for a transition.
func (tc *TransitionConfig) Event(event Event) *TransitionConfig {
	tc.event = event
	return tc
}

// EntryAction adds an entry action for a transition.
func (tc *TransitionConfig) EntryAction(action func(ctx context.Context, args ...any)) *TransitionConfig {
	tc.entryAction = action
	return tc
}

// ExitAction adds an exit action for a transition.
func (tc *TransitionConfig) ExitAction(action func(ctx context.Context, args ...any)) *TransitionConfig {
	tc.exitAction = action
	return tc
}

// EnterAction adds an enter action for a transition.
func (tc *TransitionConfig) EnterAction(action func(ctx context.Context, args ...any)) *TransitionConfig {
	tc.enterAction = action
	return tc
}

// Guard adds a guard condition to the transition.
func (tc *TransitionConfig) Guard(guard func(ctx context.Context, arg ...any) bool) *TransitionConfig {
	tc.guard = guard
	return tc
}

// EntryAction adds an entry action for a state.
func (sc *StateConfig) EntryAction(action func(ctx context.Context, args ...any)) *StateConfig {
	sc.entryAction = action
	return sc
}

// EnterAction adds an action that is triggered when entering a state.
func (sc *StateConfig) EnterAction(action func(ctx context.Context, args ...any)) *StateConfig {
	sc.enterAction = action
	return sc
}

// ExitAction adds an exit action for a state.
func (sc *StateConfig) ExitAction(action func(ctx context.Context, args ...any)) *StateConfig {
	sc.exitAction = action
	return sc
}

// Guard adds a guard condition to the state.
func (sc *StateConfig) Guard(guard func(ctx context.Context) bool) *StateConfig {
	sc.guard = guard
	return sc
}