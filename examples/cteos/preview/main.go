package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjE0LCJ1c3IiOjgsInRwIjoyLCJpYXQiOjE2NzIyNTAzMzV9.TY8-eAg6gUFSo55epFL-UoPTD3XAUJMl8SxUcAsCr4o"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	cteos := sdk_cloud_dfe.Cteos(config)

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo.")
	}

	payload := map[string]interface{}{
		"cfop":                    "5353",
		"natureza_operacao":       "PRESTACAO DE SERVICO",
		"numero":                  "64",
		"serie":                   "1",
		"data_emissao":            "2020-11-24T03:00:00-03:00",
		"tipo_operacao":           "0",
		"codigo_municipio_envio":  "2408003",
		"nome_municipio_envio":    "MOSSORO",
		"uf_envio":                "RN",
		"tipo_servico":            "6",
		"codigo_municipio_inicio": "2408003",
		"nome_municipio_inicio":   "Mossoro",
		"uf_inicio":               "RN",
		"codigo_municipio_fim":    "2408003",
		"nome_municipio_fim":      "Mossoro",
		"uf_fim":                  "RN",
		"valores": map[string]interface{}{
			"servico":       "0.00",
			"valor_total":   "0.00",
			"valor_receber": "0.00",
			"quantidade":    "10.00",
		},
		"imposto": map[string]interface{}{
			"icms": map[string]interface{}{
				"situacao_tributaria":           "99",
				"valor_base_calculo":            "0.00",
				"aliquota":                      "12.00",
				"valor":                         "0.00",
				"aliquota_reducao_base_calculo": "50.00",
			},
			"federais": map[string]interface{}{
				"valor_pis":    "0.00",
				"valor_cofins": "0.00",
				"valor_ir":     "12.00",
				"valor_inss":   "0.00",
				"valor_csll":   "50.00",
			},
		},
		"modal_rodoviario": map[string]interface{}{
			"taf":                      "020335171251",
			"numero_registro_estadual": "0203351712510203351712515",
		},
		"tomador": map[string]interface{}{
			"indicador_inscricao_estadual": "9",
			"cpf":                          "01234567890",
			"inscricao_estadual":           nil,
			"nome":                         "EMPRESA MODELO",
			"razao_social":                 "EMPRESA MODELO",
			"telefone":                     "8499995555",
			"endereco": map[string]interface{}{
				"logradouro":       "AVENIDA TESTE",
				"numero":           "444",
				"bairro":           "CENTRO",
				"codigo_municipio": "2408003",
				"nome_municipio":   "Mossoro",
				"cep":              "59603330",
				"uf":               "RN",
				"codigo_pais":      "1058",
				"nome_pais":        "BRASIL",
				"email":            "teste@teste.com.br",
			},
		},
		"componentes_valor": []map[string]interface{}{
			{
				"nome":  "teste2",
				"valor": "1999.00",
			},
		},
		"observacao": "",
	}

	resp, err := cteos.Preview(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
