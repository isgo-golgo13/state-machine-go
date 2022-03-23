package svckit

import (
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

