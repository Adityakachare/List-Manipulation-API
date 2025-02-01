package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManipulateList(t *testing.T) {
	// Test when list is empty and input is 5
	intList = []int{}
	manipulateList(5)
	assert.Equal(t, []int{5}, intList)

	// Test when list is [5] and input is 10
	intList = []int{5}
	manipulateList(10)
	assert.Equal(t, []int{5, 10}, intList)

	// Test when list is [5, 10] and input is -6
	intList = []int{5, 10}
	manipulateList(-6)
	assert.Equal(t, []int{9}, intList)

	// Test when list is [9] and input is 100
	intList = []int{9}
	manipulateList(100)
	assert.Equal(t, []int{9, 100}, intList)

	// Test when list is [9, 100] and input is -90
	intList = []int{9, 100}
	manipulateList(-90)
	assert.Equal(t, []int{19}, intList)
}
