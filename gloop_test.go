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
        want State
    }{
        {
            State{0, [][]int{}, map[string]bool{}},
            State{
                4,
                [][]int{
                    {1,2,3},
                    {4,5,6},
                },
                map[string]bool{"wookies": true},
            },
            State{
                1,
                [][]int{
                    {1,2,3},
                    {4,5,6},
                },
                map[string]bool{"wookies": true},
            },
        },
    }
    for n, tt := range tests {
        m := fmt.Sprintf("Epoch:test= %d %d", tt.state.Epoch, n)
        t.Run(m, func(t *testing.T) {
            if err := tt.state.update(&tt.input); err != nil {
                t.Errorf("update error: %s", err)
            }
            if diff := cmp.Diff(tt.state, tt.want); diff != "" {
                t.Errorf("update() mismatch (-want +got):\n%s", diff)
            }
        })
    }
}
