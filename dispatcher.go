package calculations

import "fmt"

//коллекция методик расчета + методы управления ими
type Dispatcher struct {
	methods map[string]interface{}
}

func (d *Dispatcher) initMethodsMap() {
	if d.methods == nil {
		d.methods = make(map[string]interface{})
	}
}

func (d *Dispatcher) Register(obj interface{}) bool {
	method := toMethodHeader(obj)
	if method == nil {
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

func (d *Dispatcher) Count() int{
	return len(d.methods)
}

func (d *Dispatcher) GetMethods() []MethodSpec {
	result := make([]MethodSpec, len(d.methods))

	n := int(0)
	for key, value := range d.methods {
		result[n] = MethodSpec{
			Id:key,
			Name:value.(MethodHeader).GetName(),
			Parameters:getParameters(value),
		}
		n++
	}

	return result
}

func (d *Dispatcher) Calculate(id string, values map[string]float64) (float64, error) {
	obj, ok := d.methods[id]
	if !ok {
		return 0, fmt.Errorf("Методика '%v' не зарегистрирована", id)
	}

	//смотрим, может это вообще посчитать нельзя
	calculator := toMethodSingleValue(obj)
	if calculator == nil {
		return 0, fmt.Errorf("Методика '%v' не реализует вычисление (MethodSingleValue)", id)
	}

	//сверяем параметры методики с тем, что передано
	parameters := make(map[string]float64)
	for _, p_name := range toMethodHeader(obj).GetParameters() {
		value, ok := values[p_name]
		if !ok {
			return 0, fmt.Errorf("Не указано значение для параметра '%v'", p_name)
		}
		parameters[p_name] = value
	}

	return calculator.Calculate(parameters), nil
}

func getParameters(obj interface{}) []ParameterSpec {
	method := toMethodHeader(obj)
	if method == nil {
		return nil
	}

	names := method.GetParameters()
	spec := make([]ParameterSpec, len(names))

	for idx, ee := range names {
		spec[idx] = ParameterSpec{Id:ee, Name:getParameterName(obj,ee)}
	}

	return spec
}
