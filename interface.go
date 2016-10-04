package calculations

import "github.com/andrushk/calculations/results"

//если тип, хочет называться методикой и регистрироваться в Dispatcher,
//то он должен реализовывать этот интерфейс
type Method interface {
	//каждая методика должна иметь  свое уникальное имя (Id)
	GetId() string

	//возвращает список необходимых для расчета параметров
	GetParameters() []string
}

//Dispatcher может попросить у методики описания ее параметров (и описание самой методики)
//если методика не реализует данный интерфейс, то описание просто вернется пустым
type MethodDetails interface {
	//возвращает описание(для пользователя) параметра или методики
	GetDescription(id string) string
}

//Dispatcher, для выполнения расчета, будет вызывать метод Calculate методики,
//если методика не реализует данный интерфейс то при расчете произойдет ошибка
type MethodSingleValue interface {
	Calculate(values map[string]float64) results.MethodResult
}