package romanumeral

// IRomanNumeralGenerator is the interface for converting
// Arabic numerals to Roman numerals
type IRomanNumeralGenerator interface {
	Generate(number int) (string, error)
}

// RomanNumeralGenerator is an implementation of IRomanNumeralGenerator for converting
// Arabic numerals to Roman numerals
type RomanNumeralGenerator struct {
	symbols []SymbolPair // This is a slice not a map because it needs to be ordered, largest number first
}

// NewRomanNumeralGenerator initializes and returns a new RomanNumeralGenerator
func NewRomanNumeralGenerator() *RomanNumeralGenerator {
	generator := RomanNumeralGenerator{}
	pairs := []SymbolPair{
		SymbolPair{Arabic: 1000, Roman: "M"},
		SymbolPair{Arabic: 500, Roman: "D"},
		SymbolPair{Arabic: 100, Roman: "C"},
		SymbolPair{Arabic: 50, Roman: "L"},
		SymbolPair{Arabic: 10, Roman: "X"},
		SymbolPair{Arabic: 5, Roman: "V"},
		SymbolPair{Arabic: 1, Roman: "I"},
	}

	generator.symbols = pairs

	return &generator
}

// Generate a string Roman numeral from an integer number
func (generator *RomanNumeralGenerator) Generate(number int) (string, error) {
	roman := ""
	for _, pair := range generator.symbols {
		if pair.Arabic == number {
			roman = pair.Roman
		}
	}
	return roman, nil
}
