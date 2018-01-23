package romanumeral_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chvck/romanumeral"
)

func TestNewSymbolPair(t *testing.T) {
	pair := romanumeral.NewSymbolPair(50, "L")

	assert.Equal(t, "L", pair.Roman)
	assert.Equal(t, 50, pair.Arabic)
}

func TestNewSymbolPair2(t *testing.T) {
	pair := romanumeral.NewSymbolPair(100, "C")

	assert.Equal(t, "C", pair.Roman)
	assert.Equal(t, 100, pair.Arabic)
}
