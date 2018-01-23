package romanumeral

// SymbolPair represents the relationship between an
// Arabic numeral and a Roman numeral
type SymbolPair struct {
	Arabic int
	Roman  string
}

// NewSymbolPair initializes and returns a new SymbolPair
func NewSymbolPair(arabic int, roman string) *SymbolPair {
	return &SymbolPair{
		Arabic: arabic,
		Roman:  roman,
	}
}
