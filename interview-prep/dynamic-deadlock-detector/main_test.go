package main

import (
	"testing"

	_assert "github.com/stretchr/testify/assert"
)

func Test_Process(t *testing.T) {

	tests := []struct {
		name string

		event *Event

		inGraph       func() [][]int
		inResourceIdx map[string]string

		expectedCycle    []string
		expectedDeadlock bool
		outGraph         func() [][]int
		outResourceIdx   map[string]string
	}{
		{
			name: "acquire event - requested GPU is available",

			event: &Event{
				Kind:     "acquire",
				Thread:   "T21",
				Resource: "gpu0",
			},

			inGraph:       func() [][]int { return [][]int{} },
			inResourceIdx: map[string]string{},

			expectedCycle:    nil,
			expectedDeadlock: false,
			outGraph:         func() [][]int { return [][]int{} },
			outResourceIdx:   map[string]string{"gpu0": "T21"},
		},
		{
			name: "acquire event - requested GPU is in use",

			event: &Event{
				Kind:     "acquire",
				Thread:   "T21",
				Resource: "gpu0",
			},

			inGraph: func() [][]int { return make([][]int, 100) },
			inResourceIdx: map[string]string{
				"gpu0": "T99",
			},

			expectedCycle:    nil,
			expectedDeadlock: false,
			outGraph: func() [][]int {
				g := make([][]int, 100)
				g[21] = []int{99}
				return g
			},
			outResourceIdx: map[string]string{"gpu0": "T99"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := _assert.New(t)

			assert.Equal(0, 0)

			inGraph := tt.inGraph()
			outGraph := tt.outGraph()

			d := NewDetector(
				WithGraph(inGraph),
				WithResourceIdx(tt.inResourceIdx),
			)
			cycle, deadlocked := d.Process(tt.event)

			assert.Equal(tt.expectedCycle, cycle)
			assert.Equal(tt.expectedDeadlock, deadlocked)
			assert.Equal(outGraph, inGraph)
			assert.Equal(tt.outResourceIdx, tt.inResourceIdx)
		})
	}
}
