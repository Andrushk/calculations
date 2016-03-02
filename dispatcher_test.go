package calculations

import (
	"testing"
	"github.com/andrushk/calculations/mock"
)

func TestDispatcherAddMethod(t *testing.T){
	ok := (&Dispatcher{}).Register(&mock.MethodSum{})

	if !ok{
		t.Fatal("Не удалось зарегист")
	}
}
