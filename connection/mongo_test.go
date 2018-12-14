package connection_test

import (
	"testing"

	"github.com/bickyeric/coverage-test/connection"
	"github.com/stretchr/testify/assert"
)

func TestMongo(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
	assert.NotPanics(t, func() { connection.MongoSession() }, "Eh malah error")
}
