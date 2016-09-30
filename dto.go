package calculations

//описание методики, которое формирует диспетчер при запросе списка методик
type MethodSpec struct {
	Id   string
	Name string
	Parameters []ParameterSpec
}

//описание одного параметра методики
type ParameterSpec struct {
	Id   string
	Name string
}

//name - имя параметра, value - расчитанное значение
type Detail struct {
	Name string `json:"name"`
	Value float64 `json:"value"`
}

func NewDetail(name string, value float64) *Detail {
	d := Detail{}
	d.Name = name
	d.Value = value
	return &d
}