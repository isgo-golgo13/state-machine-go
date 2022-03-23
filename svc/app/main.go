package main

import (
	"fmt"

	sk "github.com/isgo-golgo13/state-machine-go/svckit"
)

func main() {

	stateMachine := sk.New()

	initState := stateMachine.Init("locked")
	unlockedSate := stateMachine.NewState("unlocked")

	coinRule := sk.NewRule(sk.Operator("eq"), sk.Event("coin"))
	pushRule := sk.NewRule(sk.Operator("eq"), sk.Event("push"))

	stateMachine.LinkStates(initState, unlockedSate, coinRule)
	stateMachine.LinkStates(unlockedSate, initState, pushRule)

	stateMachine.LinkStates(initState, initState, pushRule)
	stateMachine.LinkStates(unlockedSate, unlockedSate, coinRule)

	fmt.Printf("Starting state is --------------> %s\n", stateMachine.CurrentState.String())

	events := []string{"coin", "push"}
	stateMachine.Compute(events, true)

	fmt.Printf("Finishing state is --------------> %s\n", stateMachine.CurrentState.String())

}
