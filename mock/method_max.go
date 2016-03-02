package mock

//находит максимум двух чисел
type MethodMax struct {
}

func (m *MethodMax) GetId() string {
	return "max"
}

func (m *MethodMax) GetName() string {
	return "Calculate Max(x,y)"
}

func (m *MethodMax) GetParameters() []string {
	return nil
}
