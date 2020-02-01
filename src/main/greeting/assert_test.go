package greeting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAssertion(t *testing.T) {
	assert.Equal(t, 123, 123, "They should be equal!")

	var request string = ""
	assert.Empty(t, request, "%s should be empty.", request)
}
