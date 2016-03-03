package calculations

//привести interface{} к методике
func toMethodHeader(obj interface{}) MethodHeader {
	method, ok := obj.(MethodHeader)

	if !ok {
		//хорошо бы еще обругаться, что попытка зарегистрировать что-то не являющееся методикой
		return nil
	}

	return method
}

//если методика реализует MethodParameters, то получить имя параметра по его Id
func getParameterName(source interface{}, id string) string {
	m, ok := source.(MethodParameters)
	if !ok {
		return ""
	}

	return m.GetDescription(id)
}