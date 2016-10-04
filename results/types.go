package results

type NumberList map[string] float64

// простое число
type Number float64

func (n Number) GetTotal() float64 {
	return float64(n)
}

func (n Number) GetDetails() NumberList {
	return nil
}


// сумма с деталями
type NumberSum NumberList

func (ns NumberSum) GetTotal() float64 {
	var sum float64 = 0
	for _, value := range ns {
		sum += value
	}
	return sum
}

func (ns NumberSum) GetDetails() NumberList {
	return NumberList(ns)
}
