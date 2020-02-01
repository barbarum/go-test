package greeting_test

import "testing"
import "github.com/stretchr/testify/mock"

type MyMockObject struct {
	mock.Mock
}

func (m *MyMockObject) testFunction(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}
func TestMock(t *testing.T) {
	testObj := new(MyMockObject)
	testObj.On("testFunction", 123).Return(true, nil)
}
