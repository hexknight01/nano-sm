package nanosm

import "context"

// StateConfig represents State's identity and set of handlers associated with it
type StateConfig struct {
	sm            *StateMachine
	state         State
	stateHandlers StateHandlers
}

type StateHandlers struct {
	entryAction []ActionFunc
	exitAction  []ActionFunc
	enterAction []ActionFunc
	guard       []func(ctx context.Context) bool
}

// Event sets the event for a transition.
// func (tc *StateConfig) State(state State) *StateConfig {
// 	if _, ok := tc.sm.states[state]; !ok {
// 		panic("this state is already configured")
// 	}
// 	tc.initConfig(state)
// 	return tc
// }

// func (tc *StateConfig) initConfig(state State) *StateConfig {
// 	tc.sm.states[state] = &StateConfig{}
// 	return tc
// }

// EntryAction adds an entry action for a state.
func (sc *StateConfig) EntryAction(action func(ctx context.Context, args ...any) error) *StateConfig {
	sc.stateHandlers.entryAction = append(
		sc.stateHandlers.entryAction,
		action,
	)
	return sc
}

// EnterAction adds an action that is triggered when entering a state.
func (sc *StateConfig) EnterAction(action func(ctx context.Context, args ...any) error) *StateConfig {
	sc.stateHandlers.enterAction = append(
		sc.stateHandlers.enterAction,
		action,
	)
	return sc
}

// ExitAction adds an exit action for a state.
func (sc *StateConfig) ExitAction(action func(ctx context.Context, args ...any) error) *StateConfig {
	sc.stateHandlers.exitAction = append(
		sc.stateHandlers.exitAction,
		action,
	)
	return sc
}

// Guard adds a guard condition to the state.
func (sc *StateConfig) Guard(guard func(ctx context.Context) bool) *StateConfig {
	sc.stateHandlers.guard = append(sc.stateHandlers.guard, guard)
	return sc
}

func (sc *StateConfig) Build() *StateMachine {
	return sc.sm
}
