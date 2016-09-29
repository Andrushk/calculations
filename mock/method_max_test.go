package mock

import (
	"testing"
	"github.com/andrushk/calculations"
)

func TestMaxShouldHaveTwoParameters(t *testing.T) {
	if count := len((&MethodMax{}).GetParameters()); count != 2 {
		t.Fatalf("MethodMax должен содержать 2 параметра, а содержит %v", count)
	}
}

func TestMaxShouldImplementMethodHeader(t *testing.T) {
	var method interface{} = &MethodMax{}
	_, ok := method.(calculations.Method)

	if !ok {
		t.Fatal("MethodMax в целях тестирования должен реализовывать Method")
	}
}

func TestMaxNotImplementMethodParameters(t *testing.T) {
	var method interface{} = &MethodMax{}
	_, ok := method.(calculations.MethodDescription)

	if ok {
		t.Fatal("MethodMax в целях тестирования не должен реализовывать MethodDescription")
	}
}