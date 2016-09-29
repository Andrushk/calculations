package mock

import (
	"testing"
	"github.com/andrushk/calculations"
)

func TestSumShouldHaveTwoParameters(t *testing.T) {
	if count := len((&MethodMax{}).GetParameters()); count != 2 {
		t.Fatalf("MethodSum должен содержать 2 параметра, а содержит %v", count)
	}
}

func TestSumShouldImplementMethodHeader(t *testing.T) {
	var method interface{} = &MethodSum{}
	_, ok := method.(calculations.Method)

	if !ok {
		t.Fatal("MethodSum в целях тестирования должен реализовывать Method")
	}
}

func TestSumShouldImplementMethodParameters(t *testing.T) {
	var method interface{} = &MethodSum{}
	_, ok := method.(calculations.MethodDescription)

	if !ok {
		t.Fatal("MethodSum в целях тестирования должен реализовывать MethodDetails")
	}
}
