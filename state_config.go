package nanosm

import "context"

// StateConfig represents State's identity and set of handlers associated with it
type StateConfig struct {
	sm            *StateMachine
	stateHandlers map[State]*StateHandlers
}

type StateHandlers struct {
	state       State
	entryAction ActionFunc
	exitAction  ActionFunc
	enterAction ActionFunc
	guard       func(ctx context.Context) bool
}

// Event sets the event for a transition.
func (tc *StateConfig) State(state State) *StateConfig {
	if _, ok := tc.sm.states[state]; !ok {
		panic("this state is already configured")
	}
	tc.initConfig(state)
	return tc
}

func (tc *StateConfig) initConfig(state State) *StateConfig {
	tc.sm.states[state] = &StateConfig{}
	return tc
}
