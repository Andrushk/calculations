package mock

//Возвращает число Пи
//реализует интерфейс MethodHeader
//а вот интерфейс MethodParameters не реализует, не нужен он ему (что используется в тестах)
type MethodPi struct {
}

const Mock_method_pi_id = "pi"

func (m *MethodPi) GetId() string {
	return Mock_method_pi_id
}

func (m *MethodPi) GetName() string {
	return "Return Pi"
}

func (m *MethodPi) GetParameters() []string {
	return nil
}
