package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	for i := 0; i <= 100; i++ {
		assert.Equal(t, i, Abs(i))
	}
	for i := -1; i >= -100; i-- {
		assert.Equal(t, -i, Abs(i))
	}

}