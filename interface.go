package calculations

type Method interface {
	GetId() string

	//возвращает список необходимых для расчета параметров
	GetParameters() []string
}

type MethodDetails interface {
	//возвращает описание(для пользователя) параметра или методики
	GetDescription(id string) string
}

//методика, возвращающая одно значение
type MethodSingleValue interface {
	Calculate(values map[string]float64) float64
}

//методика, возвращающая набор значений
//type MethodMultiValues interface {
//Calculate(repo Repo) Repo
//}