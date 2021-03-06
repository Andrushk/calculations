package calculations

import (
	"testing"
	"github.com/andrushk/calculations/mock"
)

func TestAssistToMethodHeader(t *testing.T) {
	var realMethod interface{} = &mock.MethodSum{}
	_, ok := (toMethodHeader(realMethod)).(Method)
	if !ok {
		t.Fatal("MethodSum не был приведен к Method")
	}

	var wrongMethod interface{} = &mock.MethodWrong{}
	_, ok = (toMethodHeader(wrongMethod)).(Method)
	if ok {
		t.Fatal("Ошибка, что MethodWrong удалось привести к Method")
	}
}

func TestAssistToMethodSingleValue(t *testing.T) {
	var realMethod interface{} = &mock.MethodSum{}
	_, ok := (toMethodSingleValue(realMethod)).(MethodSingleValue)
	if !ok {
		t.Fatal("MethodSum не был приведен к MethodSingleValue")
	}

	var wrongMethod interface{} = &mock.MethodWrong{}
	_, ok = (toMethodSingleValue(wrongMethod)).(MethodSingleValue)
	if ok {
		t.Fatal("Ошибка, что MethodWrong удалось привести к MethodSingleValue")
	}
}

func TestAssistGetParameterName(t *testing.T) {
	method := &mock.MethodSum{}
	parameters := method.GetParameters()

	if count := len(parameters); count == 0 {
		t.Fatal("Метод не имеет параметров, тест бессмысленный")
	}

	for _, p := range parameters {
		p_name := getDescription(method, p)
		if len(p_name) == 0 {
			t.Fatal("Не удалось получить имя параметра методики")
		}
	}
}

func TestAssistGetParameterNameWrong(t *testing.T) {
	//используется методика, не реализующая MethodParameters, т.е. не возвращающая имена параметров
	method := &mock.MethodMax{}
	parameters := method.GetParameters()

	if count := len(parameters); count == 0 {
		t.Fatal("Метод не имеет параметров, тест бессмысленный")
	}

	for _, p := range parameters {
		p_name := getDescription(method, p)
		if len(p_name) != 0 {
			t.Fatal("Удалось получить имя параметра, это ошибка")
		}
	}
}