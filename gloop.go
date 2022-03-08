package gloop

import "fmt"

// Loop interface for a game which has one main loop and, potentially,
// many nested loops.
type Loop interface {
	tick() error
	update(*State) error
}

// State structure holding the persistent data for a loop or a game.
type State struct {
	Epoch int
	Geom  [][]int
	Flags map[string]bool
}

func (s *State) tick() error {
	s.Epoch++
	return nil
}
func (s *State) update(st *State) error {
	if err := s.tick(); err != nil {
		return err
	}
	s.Geom = st.Geom
	s.Flags = st.Flags
	return nil
}

// Game can control a list of Loops and dispatch actions against them.
type Game interface {
	TickLoops() error
}

// StateLoop structure which inherits the functionality and storage of State, and
// holds a list of States which are themselves Loops.
type StateLoop struct {
	State
	Loops []*State
}

// TickLoops advances every epoch in Loops.
func (g *StateLoop) TickLoops() error {
	fmt.Println("pre: ", g.Loops)
	for _, s := range g.Loops {
		s.tick()
	}
	fmt.Println("post: ", g.Loops)
	return nil
}
