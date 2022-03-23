package svckit

import (
	"errors"
	"fmt"
	"strconv"

	"gonum.org/v1/gonum/graph"
)

type State struct {
	Id    int64
	Value interface{}
}

type Link struct {
	Id    int64
	T, F  graph.Node
	Rules map[Operator]Event
}

func (n State) ID() int64 {
	return n.Id
}


func (l Link) From() graph.Node {
	return l.F
}

func (l Link) To() graph.Node {
	return l.T
}

func (l Link) ID() int64 {
	return l.Id
}

func (l Link) ReversedLine() graph.Line {
	return Link{F: l.T, T: l.F}
}

func (n State) String() string {
	switch n.Value.(type) {
	case int:
		return strconv.Itoa(n.Value.(int))
	case float32:
		return fmt.Sprintf("%f", n.Value.(float32))
	case float64:
		return fmt.Sprintf("%f", n.Value.(float64))
	case bool:
		return strconv.FormatBool(n.Value.(bool))
	case string:
		return n.Value.(string)
	default:
		return ""
	}
}