package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjQ2MSwidXNyIjoxNzAsInRwIjoyLCJpYXQiOjE2NTE1MDYzMjR9.a0cOwP6BUDZAboYwMzoMjutCtFM8Ph-X4pLahZIB_V4"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	nfse := sdk_cloud_dfe.Nfse(config)

	payload := map[string]interface{}{
		"numero":       "1",
		"serie":        "0",
		"tipo":         "1",
		"status":       "1",
		"data_emissao": "2017-12-27T17:43:14-03:00",
		"tomador": map[string]interface{}{
			"cnpj":         "12345678901234",
			"cpf":          nil,
			"im":           nil,
			"razao_social": "Fake Tecnologia Ltda",
			"endereco": map[string]interface{}{
				"logradouro":       "Rua New Horizon",
				"numero":           "16",
				"complemento":      nil,
				"bairro":           "Jardim America",
				"codigo_municipio": "4119905",
				"uf":               "PR",
				"cep":              "81530945",
			},
		},
		"servico": map[string]interface{}{
			"codigo_tributacao_municipio":   "10500",
			"discriminacao":                 "Exemplo Serviço",
			"codigo_municipio":              "4119905",
			"valor_servicos":                "1.00",
			"valor_pis":                     "1.00",
			"valor_cofins":                  "1.00",
			"valor_inss":                    "1.00",
			"valor_ir":                      "1.00",
			"valor_csll":                    "1.00",
			"valor_outras":                  "1.00",
			"valor_aliquota":                "1.00",
			"valor_desconto_incondicionado": "1.00",
		},
		"intermediario": map[string]interface{}{
			"cnpj":         "12345678901234",
			"cpf":          nil,
			"im":           nil,
			"razao_social": "Fake Tecnologia Ltda",
		},
		"obra": map[string]interface{}{
			"codigo": "2222",
			"art":    "1111",
		},
	}

	resp, err := nfse.Preview(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
