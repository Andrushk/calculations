package results

//ответ методики (результат вычислений)
type MethodResult interface {
	GetTotal() float64
	GetDetails() NumberList
}
