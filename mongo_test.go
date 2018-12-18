package coveragetest_test

import (
	"testing"

	coveragetest "github.com/bickyeric/coverage-test"
	"github.com/stretchr/testify/assert"
)

func TestMongo(t *testing.T) {

	assert.Equal(t, 1, 1, "The two words should be the same.")
	assert.NotPanics(t, func() { coveragetest.MongoSession() }, "Eh malah error")
}
