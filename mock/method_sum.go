package mock

//суммирует два значения
//реализует интерфейсы: Method, MethodDescription, MethodSingleValue
type MethodSum struct {
}

const (
	MockMethodSumId = "sum"
	MockMethodParam1 = "summand1"
	MockMethodParam2 = "summand2"
)

func (m *MethodSum) GetId() string {
	return MockMethodSumId
}

func (m *MethodSum) GetParameters() []string {
	return []string{MockMethodParam1, MockMethodParam2}
}

func (m *MethodSum) GetDescription(id string) string {
	switch id {
	case MockMethodSumId:
		return "Calculate sum x+y"
	case MockMethodParam1:
		return "первое слагаемое"
	case MockMethodParam2:
		return "второе слагаемое"
	}

	return ""
}

func (m *MethodSum) Calculate(values map[string]float64) interface{} {
	return values[MockMethodParam1] + values[MockMethodParam2]
}