package dynamicarray_test

import (
	"testing"

	"github.com/nycdavid/scratch/interview-prep/dynamic-deadlock-detector/dynamicarray"
	_assert "github.com/stretchr/testify/assert"
)

func Test_Set_HappyPath(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "sets the value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := _assert.New(t)

			da := dynamicarray.New()
			da.Set(0, 15)
			val := da.Get(0)

			assert.Equal(15, val)
		})
	}
}
