package main

import (
	"encoding/json"
	"fmt"
	"time"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjQ2MSwidXNyIjoxNzAsInRwIjoyLCJpYXQiOjE2NTE1MDYzMjR9.a0cOwP6BUDZAboYwMzoMjutCtFM8Ph-X4pLahZIB_V4"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	nfe := sdk_cloud_dfe.Nfe(config)

	listaItens := map[string]interface{}{}

	payload := map[string]interface{}{
		"natureza_operacao":  "VENDA DENTRO DO ESTADO",
		"serie":              "1",
		"numero":             "101007",
		"data_emissao":       "2021-06-26T13:00:00-03:00",
		"data_entrada_saida": "2021-06-26T13:00:00-03:00",
		"tipo_operacao":      "1",
		"finalidade_emissao": "1",
		"consumidor_final":   "1",
		"presenca_comprador": "9",
		"intermediario": map[string]interface{}{
			"indicador": "0",
		},
		"notas_referenciadas": []map[string]interface{}{
			{
				"nfe": map[string]interface{}{
					"chave": "50000000000000000000000000000000000000000000",
				},
			},
		},
		"destinatario": map[string]interface{}{
			"cpf":                          "01234567890",
			"nome":                         "EMPRESA MODELO",
			"indicador_inscricao_estadual": "9",
			"inscricao_estadual":           nil,
			"endereco": map[string]interface{}{
				"logradouro":       "AVENIDA TESTE",
				"numero":           "444",
				"bairro":           "CENTRO",
				"codigo_municipio": "4108403",
				"nome_municipio":   "Mossoro",
				"uf":               "PR",
				"cep":              "59653120",
				"codigo_pais":      "1058",
				"nome_pais":        "BRASIL",
				"telefone":         "8499995555",
			},
		},
		"itens": listaItens,
		"frete": map[string]interface{}{
			"modalidade_frete": "0",
			"volumes": []map[string]interface{}{
				{
					"quantidade":   "10",
					"especie":      nil,
					"marca":        "TESTE",
					"numero":       nil,
					"peso_liquido": 500,
					"peso_bruto":   500,
				},
			},
		},
		"cobranca": map[string]interface{}{
			"fatura": map[string]interface{}{
				"numero":         "1035.00",
				"valor_original": "224.50",
				"valor_desconto": "0.00",
				"valor_liquido":  "224.50",
			},
		},
		"pagamento": map[string]interface{}{
			"formas_pagamento": []map[string]interface{}{
				{
					"meio_pagamento": "01",
					"valor":          "224.50",
				},
			},
		},
		"informacoes_adicionais_contribuinte": "PV: 3325 * Rep: DIRETO * Motorista:  * Forma Pagto: 04 DIAS * teste de observação para a nota fiscal * Valor aproximado tributos R$9,43 (4,20%) Fonte: IBPT",
		"pessoas_autorizadas": []map[string]interface{}{
			{
				"cnpj": "96256273000170",
			},
			{
				"cnpj": "80681257000195",
			},
		},
	}

	resp, err := nfe.Cria(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	if resp["sucesso"].(bool) {
		chave := resp["chave"].(string)
		time.Sleep(5 * time.Second)
		tentativa := 1

		for tentativa <= 5 {

			payload := map[string]interface{}{
				"chave": chave,
			}

			respC, err := nfe.Consulta(payload)

			if err != nil {
				print(err)
			}

			if respC["codigo"].(float64) == 5023 {

				if respC["sucesso"].(bool) {
					jsonData, err := json.Marshal(respC)
					if err != nil {
						fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
					}
					fmt.Println(string(jsonData))
					break

				} else {
					jsonData, err := json.Marshal(respC)
					if err != nil {
						fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
					}
					fmt.Println(string(jsonData))
					break
				}

			}

			time.Sleep(5 * time.Second)
			tentativa += 1

		}

	} else if resp["codigo"].(float64) == 5001 || resp["codigo"].(float64) == 5002 {
		jsonData, err := json.Marshal(resp["erros"])

		if err != nil {
			fmt.Printf("Erro ao converter mapa para Json: %v \n", err)
		}

		fmt.Println(string(jsonData))

	} else if resp["codigo"].(float64) == 5008 || resp["codigo"].(float64) >= 7000 {
		chave := resp["chave"].(string)

		payload := map[string]interface{}{
			"chave": chave,
		}

		respC, err := nfe.Consulta(payload)

		if err != nil {
			fmt.Println(err)
		}

		if respC["sucesso"].(bool) {

			if respC["codigo"].(float64) == 5023 {
				jsonData, err := json.Marshal(respC)

				if err != nil {
					fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
				}

				fmt.Println(string(jsonData))
			}
		}

	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
	}

	fmt.Println(string(jsonData))

}
