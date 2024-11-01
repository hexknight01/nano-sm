package main

import (
	"context"
	"fmt"

	sm "github.com/hexknight01/nano-sm"
)

func main() {
	stateMachine := sm.NewStateMachineWithExternalStorage("NONE", func(ctx context.Context, arg ...any) sm.State {
		// Fetch the current state from an external storage (e.g., database, API)
		return "NONE" // Replace with actual state fetching logic
	})
	stateMachine.ConfigState("NEW").
		// EntryAction(func(ctx context.Context, args ...any) { fmt.Println("Entry Action: NEW State") }).
		// ExitAction(func(ctx context.Context, args ...any) { fmt.Println("Exit Action: NEW State") }).
		// EnterAction(func(ctx context.Context, args ...any) { fmt.Println("Enter Action: NEW State") })

		// Define actions for the DEPLOYING_QC state with a guard.
		stateMachine.State("DEPLOYING_QC").
		EntryAction(func(ctx context.Context, args ...any) { fmt.Println("Entry Action: DEPLOYING_QC State") }).
		ExitAction(func(ctx context.Context, args ...any) { fmt.Println("Exit Action: DEPLOYING_QC State") }).
		EnterAction(func(ctx context.Context, args ...any) { fmt.Println("Enter Action: DEPLOYING_QC State") }).
		Guard(func(ctx context.Context) bool {
			fmt.Println("Checking guard for DEPLOYING_QC")
			return true // Guard logic
		})
	// Define a transition from NEW to DEPLOYING_QC on event TICKET_CREATION_STARTED.
	transition := stateMachine.Transition("NEW", "DEPLOYING_QC").
		Event("TICKET_CREATION_STARTED").
		EntryAction(func(ctx context.Context, args ...any) { fmt.Println("Entry Action on Transition NEW -> DEPLOYING_QC") }).
		ExitAction(func(ctx context.Context, args ...any) { fmt.Println("Exit Action on Transition NEW -> DEPLOYING_QC") }).
		Guard(func(ctx context.Context, args ...any) bool {
			fmt.Println("Guard on Transition NEW -> DEPLOYING_QC")
			return true
		})
	// Add the transition to the state machine.
	stateMachine.AddTransition(transition)

	stateMachine.TriggerEvent("TICKET_CREATION_STARTED")
}
