package function

import "testing"

func TestExampleFailed(t *testing.T) {
	str := Handle([]byte("hello"))
	t.Error(str)
}
