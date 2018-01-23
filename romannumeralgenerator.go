package romanumeral

// IRomanNumeralGenerator is the interface for converting
// Arabic numerals to Roman numerals
type IRomanNumeralGenerator interface {
	Generate(number int) (string, error)
}

// RomanNumeralGenerator is an implementation of IRomanNumeralGenerator for converting
// Arabic numerals to Roman numerals
type RomanNumeralGenerator struct {
}

// NewRomanNumeralGenerator initializes and returns a new RomanNumeralGenerator
func NewRomanNumeralGenerator() *RomanNumeralGenerator {
	return &RomanNumeralGenerator{}
}

// Generate a string Roman numeral from an integer number
func (generator *RomanNumeralGenerator) Generate(number int) (string, error) {
	return "", nil
}
