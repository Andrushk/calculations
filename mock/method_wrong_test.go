package mock

import (
	"testing"
	"github.com/andrushk/calculations"
)

func TestWrongNotImplementMethodHeader(t *testing.T) {
	var method interface{} = &MethodWrong{}
	_, ok := method.(calculations.Method)

	if ok {
		t.Fatal("MethodWrong в целях тестирования не должен реализовывать Method")
	}
}