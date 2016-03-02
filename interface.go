package calculations

type MethodHeader interface {
	GetId() string

	//возвращает понятное пользователю имя методики
	GetName() string

	//возвращает список необходимых для расчета параметров
	GetParameters() []string
}

type MethodParameters interface {
	//возвращает описание(для пользователя) параметра
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