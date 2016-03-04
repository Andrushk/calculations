package mock

//суммирует два значения
//реализует интерфейсы: MethodHeader, MethodParameters
type MethodSum struct {
}

const (
	Mock_method_sum_id = "sum"
	Mock_method_parameter1 = "summand1"
	Mock_method_parameter2 = "summand2"
)

func (m *MethodSum) GetId() string {
	return Mock_method_sum_id
}

func (m *MethodSum) GetName() string {
	return "Calculate sum x+y"
}

func (m *MethodSum) GetParameters() []string {
	return []string{Mock_method_parameter1, Mock_method_parameter2}
}

func (m *MethodSum) GetDescription(id string) string {
	switch id {
	case Mock_method_parameter1:
		return "первое слагаемое"
	case Mock_method_parameter2:
		return "второе слагаемое"
	}

	return ""
}

func (m *MethodSum) Calculate(values map[string]float64) float64 {
	return values[Mock_method_parameter1] + values[Mock_method_parameter2]
}