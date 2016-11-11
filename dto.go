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