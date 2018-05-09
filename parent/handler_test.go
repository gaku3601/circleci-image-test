package function

import (
	"fmt"
	"testing"
)

func TestExampleFailed(t *testing.T) {
	str := Handle([]byte("hello"))
	fmt.Println(str)
}
