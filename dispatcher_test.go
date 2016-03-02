package calculations

import (
	"testing"
	"github.com/andrushk/calculations/mock"
)

func TestNewDispatcherZeroCount(t *testing.T) {
	if length := (&Dispatcher{}).Count(); length != 0 {
		t.Fatalf("Кол-во методик в только что созданном диспетчере должно быть равным 0, а по факту %v", length)
	}
}

func TestDispatcherAddMethod(t *testing.T){
	ok := (&Dispatcher{}).Register(&mock.MethodSum{})

	if !ok{
		t.Fatal("Не удалось зарегистрировать методику")
	}
}

func TestDispatcherAddFakeMethod(t *testing.T){
	ok := (&Dispatcher{}).Register(&mock.MethodWrong{})

	if ok{
		t.Fatal("Диспетчер съел объект не являющийся методикой")
	}
}

func TestDispatcherAddMethodTwice(t *testing.T){
	d := &Dispatcher{}
	d.Register(&mock.MethodSum{})
	twice := d.Register(&mock.MethodSum{})

	if twice{
		t.Fatal("Одну методику получилось зарегистрировать дважды")
	}
}

func TestDispatcherAddSeveralMethod(t *testing.T) {
	d := &Dispatcher{}
	ok1 := d.Register(&mock.MethodSum{})
	ok2 := d.Register(&mock.MethodMax{})

	if count := d.Count(); !ok1 || !ok2 || count != 2 {
		t.Fatalf("Проблемы с регистрацией двух методик, кол-во зарегистрированных %v", count)
	}
}