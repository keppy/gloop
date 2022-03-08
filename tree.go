package gloop

import (
	"fmt"

	"github.com/petar/GoLLRB/llrb"
)

// StateNode structures is a state that embeds llrb.Node to make a tree
type StateNode struct {
	State
	llrb.Node
}

// Less fulfills the LLRB Item interface
func (s StateNode) Less(than llrb.Item) bool {
	return s.Epoch < than.(StateNode).Epoch
}

// GameTree can has loop tree game state
type GameTree struct {
	Game
	llrb.LLRB
}

// PourStateNode transfers llrb.Item values in to a StateNode
func (g *GameTree) PourStateNode(i llrb.Item) *StateNode {
	s := &StateNode{
		State: State{
			i.(StateNode).Epoch,
			i.(StateNode).Geom,
			i.(StateNode).Flags,
		},
	}
	return s
}

// TickLoops increments only the lowest epoch value state
func (g *GameTree) TickLoops() error {
	l := g.Min()
	n := g.PourStateNode(l)
	n.tick()
	g.DeleteMin()
	g.InsertNoReplace(n)
	fmt.Println("TickLoops() pre-state: ", l.(StateNode).Epoch)
	fmt.Println("TickLoops() post-state: ", n.Epoch)
	return nil
}
