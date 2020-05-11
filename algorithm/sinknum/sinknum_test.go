package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberSink(t *testing.T) {
	testCase := []struct {
		input []int
		expect []int
	}{
		{
			[]int{1,0,2,0,3,0,4,0,5,0},
			[]int{1,2,3,4,5,0,0,0,0,0},
		},
		{
			[]int{1,0,0,0,0},
			[]int{1,0,0,0,0},
		},
		{
			[]int{0,0,0,0,1},
			[]int{1,0,0,0,0},
		},
		{
			[]int{0,0,1,0,0},
			[]int{1,0,0,0,0},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
	}
	for _, c :=range testCase {
		got := NumberSink(c.input)
		assert.Equal(t, got, c.expect)
	}
}
