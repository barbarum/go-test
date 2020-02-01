package greeting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Set up test suite
func (suite *SampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// Run all test cases through normal Test* (t *testing.T) pattern
func TestSuite(t *testing.T) {
	suite.Run(t, new(SampleTestSuite))
}

// List all test cases, it's written and included by the above TestSuite method.
func (suite *SampleTestSuite) TestVariableSetup() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}
func (suite *SampleTestSuite) TestVariableSetup2() {
	assert.NotEqual(suite.T(), 6, suite.VariableThatShouldStartAtFive)
}
