package mock

type MethodSum struct {
}

const mockMethodName string  = "sum"

func (m *MethodSum) GetId() string {
	return mockMethodName
}

func (m *MethodSum) GetName() string {
	return "Calculate sum"
}

func (m *MethodSum) GetParameters() []string {
	return nil
}