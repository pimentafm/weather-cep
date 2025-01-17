package entity

type City struct {
	CEP         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	UF          string
	IBGE        string
	GIA         string
	DDD         string
	SIAFI       string
	Erro        string
}

type CityResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
	Erro        string `json:"erro"`
}

func NewCity(response CityResponse) *City {
	return &City{
		CEP:         response.CEP,
		Logradouro:  response.Logradouro,
		Complemento: response.Complemento,
		Unidade:     response.Unidade,
		Bairro:      response.Bairro,
		Localidade:  response.Localidade,
		UF:          response.UF,
		IBGE:        response.IBGE,
		GIA:         response.GIA,
		DDD:         response.DDD,
		SIAFI:       response.SIAFI,
		Erro:        response.Erro,
	}
}
