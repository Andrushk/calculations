package mock

//подделка под методику, не реализует интерфейс Method, используется для проверки регистрации методик
type MethodWrong struct {
}

func (m *MethodWrong) GetId() string {
	return "wrong"
}

func (m *MethodWrong) GetDescription() string {
	return "nicht"
}