package calculations

import (
	"testing"
	"github.com/andrushk/calculations/mock"
)

const method_parameters_count_wrong string = "У методики '%v' должно быть %v параметра, а реально %v"

func getDispatcherWithMethods() *Dispatcher{
	d := &Dispatcher{}
	d.Register(&mock.MethodSum{})
	d.Register(&mock.MethodMax{})
	d.Register(&mock.MethodPi{})
	return d
}

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

func TestDispatcherGetMethodsList(t *testing.T) {
	methods := getDispatcherWithMethods().GetMethods()

	if count := len(methods); count != 3 {
		t.Fatalf("Список должен содержать %v методики, а содержит %v", 3, count)
	}

	//проверим списки параметров
	for _, m := range methods {
		switch m.Id {
		case mock.Mock_method_sum_id:
			if count:=len(m.Parameters); count!=2{
				t.Fatalf(method_parameters_count_wrong, m.Name, 2, count)
			}
			continue
		case mock.Mock_method_max_id:
			if count:=len(m.Parameters); count!=2{
				t.Fatalf(method_parameters_count_wrong, m.Name, 2, count)
			}
			continue
		case mock.Mock_method_pi_id:
			if count:=len(m.Parameters); count!=0{
				t.Fatalf(method_parameters_count_wrong, m.Name, 0, count)
			}
			continue
		default:
			t.Fatal("В списке методик найдены какие-то неопознанные методики")
		}
	}
}

func TestDispatcherCalculate(t *testing.T) {
	dispatcher := getDispatcherWithMethods()

	if _, err := dispatcher.Calculate(mock.Mock_method_max_id, nil); err == nil {
		t.Fatalf("Calculate для '%v' должен приводить к ошибке, т.к. она не реализует MethodSingleValue",
			mock.Mock_method_max_id)
	}

	//правильный расчет
	if value, err := dispatcher.Calculate(mock.Mock_method_sum_id,
		map[string]float64{
			mock.Mock_method_parameter1: 37,
			mock.Mock_method_parameter2: 5,
		}); value != 42 || err != nil {
		t.Fatalf("Ошибка при расчете, ответ сервера %v, ожидалось 42", value)
	}

	//бракованный расчет: ошибка в параметрах
	if value, err := dispatcher.Calculate(mock.Mock_method_sum_id,
		map[string]float64{
			mock.Mock_method_parameter1: 37,
			"wrong": 5,
		}); value != 0 || err == nil {
		t.Fatal("Ожидалась ошибка, так как не переданы необходимые параметры")
	}
}