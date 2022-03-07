package gloop

import (
    //"fmt"
)

type Loop interface {
    tick() error
    update(*State) error
}
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
    s.Geom  = st.Geom
    s.Flags = st.Flags
    return nil
}

type game interface {
    tickLoops() error
}
type Gloop struct {
    State
    Loops []State
}
func (g Gloop) tickLoops() error {
    for _, s := range g.Loops {
        s.tick()
    }
    return nil
}
