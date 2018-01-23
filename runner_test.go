package romanumeral_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chvck/romanumeral"
)

func TestRunnerRunWhenMultipleArgsFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run(1, 2, 3)

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

func TestRunnerRunWhenStringArgFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run("one")

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

func TestRunnerRunWhenFloatArgFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run(3.14)

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}

func TestRunnerRunWhenNoArgFails(t *testing.T) {
	r := romanumeral.Runner{}

	roman, err := r.Run()

	assert.NotNil(t, err)
	assert.Nil(t, roman)
}
