package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	fmt.Println(TheEarth().String())
}
