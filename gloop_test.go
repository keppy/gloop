package gloop

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestState(t *testing.T) {
	s := State{Epoch: 0}
	s.tick()
	if s.Epoch != 1 {
		t.Errorf("Epoch = %d, want 1", s.Epoch)
	}
	s.tick()
	if s.Epoch != 2 {
		t.Errorf("Epoch = %d, want 2", s.Epoch)
	}
}

func TestUpdate(t *testing.T) {
	var tests = []struct {
		state State
		input State
		want  State
	}{
		{
			State{0, [][]int{}, map[string]bool{}},
			State{
				4,
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				map[string]bool{"wookies": true},
			},
			State{
				1,
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				map[string]bool{"wookies": true},
			},
		},
	}
	for n, tt := range tests {
		m := fmt.Sprintf("Epoch:test= %d %d", tt.state.Epoch, n)
		t.Run(m, func(t *testing.T) {
			if err := tt.state.update(&tt.input); err != nil {
				t.Errorf("update() error: %s", err)
			}
			if diff := cmp.Diff(tt.state, tt.want); diff != "" {
				t.Errorf("update() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTickLoops(t *testing.T) {
	var tests = []struct {
		have Gloop
		want Gloop
	}{
		{
			Gloop{
				State: State{0, [][]int{}, map[string]bool{}},
				Loops: []*State{
					&State{0, [][]int{}, map[string]bool{}},
					&State{1, [][]int{}, map[string]bool{}},
					&State{0, [][]int{}, map[string]bool{}},
				},
			},
			Gloop{
				State: State{0, [][]int{}, map[string]bool{}},
				Loops: []*State{
					&State{1, [][]int{}, map[string]bool{}},
					&State{2, [][]int{}, map[string]bool{}},
					&State{1, [][]int{}, map[string]bool{}},
				},
			},
		},
	}
	for n, tt := range tests {
		m := fmt.Sprintf("Epoch:test= %d %d", tt.have.Epoch, n)
		t.Run(m, func(t *testing.T) {
			fmt.Println("before: ", tt.have.Loops)
			if err := tt.have.TickLoops(); err != nil {
				t.Errorf("TickLoops() error: %s", err)
			}
			fmt.Println("after: ", tt.have.Loops)
			if diff := cmp.Diff(tt.have.Loops, tt.want.Loops); diff != "" {
				t.Errorf("TickLoops() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
