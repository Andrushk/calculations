package mock

//Возвращает число Пи
//реализует интерфейс Method
//а вот интерфейс MethodDetails не реализует, не нужен он ему (что используется в тестах)
type MethodPi struct {
}

const Mock_method_pi_id = "pi"

func (m *MethodPi) GetId() string {
	return Mock_method_pi_id
}

func (m *MethodPi) GetParameters() []string {
	return nil
}

func (m *MethodPi) Calculate(values map[string]float64) interface{} {
	return 3.14
}
