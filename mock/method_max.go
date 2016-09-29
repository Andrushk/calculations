package mock

//находит максимум двух чисел
type MethodMax struct {
}

const MockMethodMaxId = "max"

func (m *MethodMax) GetId() string {
	return MockMethodMaxId
}

func (m *MethodMax) GetParameters() []string {
	return []string{"value1", "value2"}
}

func (m *MethodMax) Calculate(values map[string]float64) interface{} {
	return values[MockMethodParam1] + values[MockMethodParam2]
}