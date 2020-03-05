package generics_test

import (
	"testing"

	"github.com/disksing/generics"
)

func TestGenerics(t *testing.T) {
	defer func() {
		if s := recover(); s == nil {
			t.Error("Go should not support generics")
		}
	}()

	generics.Enable()
}
