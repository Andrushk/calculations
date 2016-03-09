package mock

//находит максимум двух чисел
type MethodMax struct {
}

const Mock_method_max_id = "max"

func (m *MethodMax) GetId() string {
	return Mock_method_max_id
}

func (m *MethodMax) GetParameters() []string {
	return []string{"value1", "value2"}
}
