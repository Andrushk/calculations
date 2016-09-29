package calculations

//привести interface{} к методике
func toMethodHeader(obj interface{}) Method {
	method, ok := obj.(Method)
	if !ok {
		//хорошо бы еще обругаться, что попытка зарегистрировать что-то не являющееся методикой
		return nil
	}
	return method
}

//если методика реализует MethodDescription, то по id получить описание параметра/методики
func getDescription(source interface{}, id string) string {
	m, ok := source.(MethodDescription)
	if !ok {
		return ""
	}
	return m.GetDescription(id)
}