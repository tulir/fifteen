package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDigits(t *testing.T) {
	assert.Equal(t, 5, Digits(12345))
	assert.Equal(t, 1, Digits(9))
	assert.Equal(t, 2, Digits(10))
	assert.Equal(t, 2, Digits(99))
	assert.Equal(t, 3, Digits(100))
	assert.Equal(t, 6, Digits(-123456))
	assert.Equal(t, 0, Digits(0))
}