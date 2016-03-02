package calculations

type Dispatcher struct {
	methods map[string]interface{}
}

func (d *Dispatcher) initMethodsMap() {
	if d.methods == nil {
		d.methods = make(map[string]interface{})
	}
}

func (d *Dispatcher) Register(obj interface{}) bool {
	method, ok := obj.(MethodHeader)
	if !ok {
		//хорошо бы еще обругаться, что попытка зарегистрировать что-то не являющееся методикой
		return false
	}

	//второй раз не регистрируем
	_, duplicate := d.methods[method.GetId()]
	if duplicate {
		return false
	}

	d.initMethodsMap()
	d.methods[method.GetId()] = method
	return true
}

//func GetMethod(id string) MethodHeader {
//	return availableMethods[id]
//}

//func GetAllMethods() MethodsList  {
//	return availableMethods
//}
