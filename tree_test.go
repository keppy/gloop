package gloop

import (
	"fmt"
	"testing"
)

func TestStateNode(t *testing.T) {
	s := StateNode{State: State{Epoch: 0}}
	s.tick()
	if s.Epoch != 1 {
		t.Errorf("Epoch = %d, want 1", s.Epoch)
	}
	s.tick()
	if s.Epoch != 2 {
		t.Errorf("Epoch = %d, want 2", s.Epoch)
	}
}

func TestStateNodeTickLoops(t *testing.T) {
	var tests = []struct {
		mid  StateNode
		want StateNode
		high StateNode
	}{
		{
			StateNode{
				State: State{3, [][]int{}, map[string]bool{}},
			},
			StateNode{
				State: State{1, [][]int{}, map[string]bool{}},
			},
			StateNode{
				State: State{5, [][]int{}, map[string]bool{}},
			},
		},
	}

	for _, tt := range tests {
		g := &GameTree{}
		g.InsertNoReplace(tt.mid)
		g.InsertNoReplace(tt.high)
		g.InsertNoReplace(tt.want)
		g.TickLoops()
		m := g.Min()
		me := m.(*StateNode).Epoch
		if me != 2 {
			t.Errorf("Min Epoch = %d, should be 2", me)
		}
		fmt.Println("min: ", m.(*StateNode).Epoch)
	}
}
