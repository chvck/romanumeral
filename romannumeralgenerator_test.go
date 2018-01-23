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

// When 4 is provided then IV should be the value
func TestGenerateWhen4ThenIV(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(4)

	assert.Nil(t, err)
	assert.Equal(t, "IV", rNumeral)
}

// When 9 is provided then IX should be the value
func TestGenerateWhen9ThenIX(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(9)

	assert.Nil(t, err)
	assert.Equal(t, "IX", rNumeral)
}

// When 40 is provided then XL should be the value
func TestGenerateWhen40ThenXL(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(40)

	assert.Nil(t, err)
	assert.Equal(t, "XL", rNumeral)
}

// When 90 is provided then XC should be the value
func TestGenerateWhen90ThenXC(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(90)

	assert.Nil(t, err)
	assert.Equal(t, "XC", rNumeral)
}

// When 400 is provided then CD should be the value
func TestGenerateWhen400ThenCD(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(400)

	assert.Nil(t, err)
	assert.Equal(t, "CD", rNumeral)
}

// When 900 is provided then CM should be the value
func TestGenerateWhen900ThenCM(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(900)

	assert.Nil(t, err)
	assert.Equal(t, "CM", rNumeral)
}

// When 20 is provided then XX should be the value
func TestGenerateWhen20ThenXX(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(20)

	assert.Nil(t, err)
	assert.Equal(t, "XX", rNumeral)
}

// When 3999 is provided then MMMCMXCIX should be the value
func TestGenerateWhen3999ThenMMMCMXCIX(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(3999)

	assert.Nil(t, err)
	assert.Equal(t, "MMMCMXCIX", rNumeral)
}

// When 73 is provided then LXXIII should be the value
func TestGenerateWhen73ThenLXXIII(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(73)

	assert.Nil(t, err)
	assert.Equal(t, "LXXIII", rNumeral)
}

// When 459 is provided then CDLIX should be the value
func TestGenerateWhen459ThenCDLIX(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(459)

	assert.Nil(t, err)
	assert.Equal(t, "CDLIX", rNumeral)
}

// When 1381 is provided then MCCCLXXXI should be the value
func TestGenerateWhen1381ThenMCCCLXXXI(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(1381)

	assert.Nil(t, err)
	assert.Equal(t, "MCCCLXXXI", rNumeral)
}

// When 2222 is provided then MMXXII should be the value
func TestGenerateWhen2222ThenMMXXII(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(2222)

	assert.Nil(t, err)
	assert.Equal(t, "MMCCXXII", rNumeral)
}

// When 4000 is provided then error should be returned
func TestGenerateWhen4000ThenError(t *testing.T) {
	generator := romanumeral.NewRomanNumeralGenerator()

	rNumeral, err := generator.Generate(4000)

	assert.NotNil(t, err)
	assert.Equal(t, "", rNumeral)
}
