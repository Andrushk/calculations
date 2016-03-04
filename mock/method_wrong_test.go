package mock

import (
	"testing"
	"github.com/andrushk/calculations"
)

func TestWrongNotImplementMethodHeader(t *testing.T) {
	var method interface{} = &MethodWrong{}
	_, ok := method.(calculations.MethodHeader)

	if ok {
		t.Fatal("MethodWrong в целях тестирования не должен реализовывать MethodHeader")
	}
}

func TestWrongNotImplementMethodSingleValue(t *testing.T) {
	var method interface{} = &MethodWrong{}
	_, ok := method.(calculations.MethodSingleValue)

	if ok {
		t.Fatal("MethodWrong в целях тестирования не должен реализовывать MethodSingleValue")
	}
}