package svckit

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
)

type Event string
type Operator string

var NodeIDCntr = 0
var LineIDCntr = 1

type StateMachine struct {
	CurrentState State 
	Graph *multi.DirectedGraph
}

func New() *StateMachine {
	s := &StateMachine {}
	s.Graph = multi.NewDirectedGraph()

	return s
}

func (s *StateMachine) Init(initStateValue interface{}) State {
	s.CurrentState = State{Id: int64(NodeIDCntr), Value: initStateValue}
	s.Graph.AddNode(s.CurrentState)
	NodeIDCntr++
	return s.CurrentState
}

func (s *StateMachine) NewState(stateValue interface{}) State {
	state := State{Id: int64(NodeIDCntr), Value: stateValue}
	s.Graph.AddNode(state)
	NodeIDCntr++
	return state
}

func (s *StateMachine) LinkStates(s1, s2 State, rule map[Operator]Event) {
	s.Graph.SetLine(Link{F: s1, T: s2, Id: int64(LineIDCntr), Rules: rule})
	LineIDCntr++
}


func (s *StateMachine) FireEvent(e Event) error {
	currentNode := s.CurrentState

	it := s.Graph.From(currentNode.Id)

	for it.Next() {
		n := s.Graph.Node(it.Node().ID()).(State)
		line := graph.LinesOf(s.Graph.Lines(currentNode.Id, n.Id))[0].(Link) // There can be one defined path between two distinct states

		for key, val := range line.Rules {
			k := string(key)
			switch k {
			case "eq":
				if val == e {
					s.CurrentState = n
					return nil
				}
			default:
				fmt.Printf("Sorry, the comparison operator '%s' is not supported\n", k)
				return errors.New("UNSUPPORTED_COMPARISON_OPERATOR")
			}
		}
	}
	return nil
}

func (s *StateMachine) Compute(events []string, printState bool) State {
	for _, e := range events {
		s.FireEvent(Event(e))
		if printState {
			fmt.Printf("%s\n", s.CurrentState.String())
		}
	}
	return s.CurrentState
}

