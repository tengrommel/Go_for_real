package visitor

import "testing"

type TestHelper struct {
	Received string
}

func (t *TestHelper)Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T)  {
	TestHelper := &TestHelper{}
	visitor := &MessageVistor{}

	t.Run("MessageA test", func(t *testing.T) {
		msg := MessageA{

		}
	})

}