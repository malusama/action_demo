package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	x int
	y int
}

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		args   args
		result int
	}{
		// TODO: Add test cases.
		{
			name: "add 1 and 2",
			args: args{
				x: 1,
				y: 2,
			},
			result: 3,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, func() int {
				t.Parallel()
				return Add(tt.args.x, tt.args.y)
			}())
		})
	}
}
