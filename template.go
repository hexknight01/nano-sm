package nanosm

import (
	"context"
	"fmt"
)

type ActionFunc func(ctx context.Context, args ...any) error
type StateTemplate struct {
	State       State
	EntryAction ActionFunc
	EnterAction ActionFunc
	ExitAction  ActionFunc
}

// ToStateTemplate returns a StateTemplate for the given state.
// If the state is not found in the StateMachine'sm states map, it returns a StateTemplate
// for the current state of the StateMachine.
//
// Parameters:
// - state: The state for which to retrieve the StateTemplate.
//
// Returns:
//   - StateTemplate: A StateTemplate containing the state, entryAction, enterAction, and exitAction
//     for the given state. If the state is not found, it returns a StateTemplate for the current state.
func (sm StateMachine) ToStateTemplate(state State) StateTemplate {
	// TODO: Double checking this critical section for concurrency edge cases
	if stateConfig, ok := sm.states[state]; ok {
		return StateTemplate{
			State:       state,
			EntryAction: stateConfig.entryAction,
			EnterAction: stateConfig.enterAction,
			ExitAction:  stateConfig.exitAction,
		}
	}
	return StateTemplate{
		State: sm.currentState,
	}
}

func (sm *StateTemplate) ExecuteEntryAction(ctx context.Context, args ...any) {
	err := sm.EntryAction(ctx, args)
	if err != nil {
		fmt.Println(err)
	}
}

func (sm *StateTemplate) ExecuteExitAction(ctx context.Context, args ...any) {
	err := sm.ExitAction(ctx, args)
	if err != nil {
		fmt.Println(err)
	}
}

func (sm *StateTemplate) ExecuteEnterAction(ctx context.Context, args ...any) {
	err := sm.EnterAction(ctx, args)
	if err != nil {
		fmt.Println(err)
	}
}
