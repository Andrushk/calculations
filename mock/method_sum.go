package mock

//суммирует две строки
type MethodSum struct {
}

func (m *MethodSum) GetId() string {
	return "sum"
}

func (m *MethodSum) GetName() string {
	return "Calculate sum x+y"
}

func (m *MethodSum) GetParameters() []string {
	return nil
}