package romanumeral_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chvck/romanumeral"
)

// When 1 is provided then I should be the value
func TestGenerateWhen1ThenI(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(1)

	assert.Nil(t, err)
	assert.Equal(t, "I", rNumeral)
}

// When 5 is provided then V should be the value
func TestGenerateWhen5ThenV(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(5)

	assert.Nil(t, err)
	assert.Equal(t, "V", rNumeral)
}

// When 10 is provided then X should be the value
func TestGenerateWhen10ThenX(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(10)

	assert.Nil(t, err)
	assert.Equal(t, "X", rNumeral)
}

// When 50 is provided then L should be the value
func TestGenerateWhen50ThenL(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(50)

	assert.Nil(t, err)
	assert.Equal(t, "L", rNumeral)
}

// When 100 is provided then C should be the value
func TestGenerateWhen100ThenC(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(100)

	assert.Nil(t, err)
	assert.Equal(t, "C", rNumeral)
}

// When 500 is provided then D should be the value
func TestGenerateWhen500ThenD(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(500)

	assert.Nil(t, err)
	assert.Equal(t, "D", rNumeral)
}

// When 1000 is provided then M should be the value
func TestGenerateWhen1000ThenM(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(1000)

	assert.Nil(t, err)
	assert.Equal(t, "M", rNumeral)
}
