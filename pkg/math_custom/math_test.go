package math_custom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountFibonacci(t *testing.T) {
	fibStorage := map[uint64]uint64{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
	}

	for num, expected := range fibStorage {
		real := CountFibonacci(num)
		assert.Equal(t, real, expected)
	}
}
