package romanumeral_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chvck/romanumeral"
)

const retRoman = "V"

type GeneratorTest struct {
}

func (g *GeneratorTest) Generate(number int) (string, error) {
	return retRoman, nil
}

// When run is called with multiple args then error should be not nil
func TestRunnerRunWhenMultipleArgsFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run(1, 2, 3)

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

// When run is called with a string arg then error should be not nil
func TestRunnerRunWhenStringArgFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run("one")

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

// When run is called with a float arg then error should be not nil
func TestRunnerRunWhenFloatArgFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run(3.14)

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

// When run is called with no arg then error should be not nil
func TestRunnerRunWhenNoArgFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run()

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

// When run is called correctly then the correct value should be returned
func TestRunnerRunWhenValidArgReturnsCorrectly(t *testing.T) {
	r := romanumeral.Runner{}
	r.Generator = &GeneratorTest{}

	roman, err := r.Run(5)

	assert.Nil(t, err)
	assert.Equal(t, retRoman, *roman)
}
