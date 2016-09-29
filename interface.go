package calculations

//если тип, хочет называться методикой и регистрироваться в Dispatcher,
//то он должен реализовывать этот интерфейс
type Method interface {
	//каждая методика должна иметь  свое уникальное имя (Id)
	GetId() string

	//возвращает список необходимых для расчета параметров
	GetParameters() []string

	//интерфейс для расчета суммы
	Calculate(values map[string]float64) interface{}
}

//Dispatcher может попросить у методики описания ее параметров (и описание самой методики)
//если методика не реализует данный интерфейс, то описание просто вернется пустым
type MethodDescription interface {
	//возвращает описание(для пользователя) параметра или методики
	GetDescription(id string) string
}