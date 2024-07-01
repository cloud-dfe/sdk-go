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

	nfcom := sdk_cloud_dfe.Nfcom(config)

	listaItens := []map[string]interface{}{
		{
			"numero_item":            "1",
			"codigo_produto":         "123",
			"descricao":              "LP 1MB",
			"codigo_classificacao":   "0400401",
			"cfop":                   "5301",
			"unidade_medida":         "1",
			"quantidade":             "1",
			"valor_unitario":         "10.00",
			"valor_desconto":         "0",
			"valor_outras_despesas":  "0",
			"valor_bruto":            "10.00",
			"indicador_devolucao":    "0",
			"informacoes_adicionais": "teste",
			"imposto": map[string]interface{}{
				"icms": map[string]interface{}{
					"situacao_tributaria": "00",
					"valor_base_calculo":  "10.00",
					"aliquota":            "18.00",
					"valor":               "1.80",
				},
				"fcp": map[string]interface{}{
					"aliquota": nil,
					"valor":    nil,
				},
			},
		},
	}

	payload := map[string]interface{}{
		"numero":                      "3",
		"serie":                       "1",
		"data_emissao":                "2024-06-20T13:23:00-03:00",
		"finalidade_emissao":          "0",
		"tipo_faturamento":            "0",
		"indicador_prepago":           "0",
		"indicador_cessao_meios_rede": "0",
		"destinatario": map[string]interface{}{
			"nome":                         "HELIO WOLFF",
			"cpf":                          "06844990960",
			"cnpj":                         "",
			"id_outros":                    "",
			"inscricao_estadual":           nil,
			"indicador_inscricao_estadual": "9",
			"endereco": map[string]interface{}{
				"logradouro":       "LOJA",
				"complemento":      nil,
				"numero":           "SN",
				"bairro":           "BANANAL",
				"codigo_municipio": "4314035",
				"nome_municipio":   "Pareci Novo",
				"uf":               "RS",
				"codigo_pais":      "1058",
				"nome_pais":        "Brasil",
				"cep":              "95783000",
				"telefone":         nil,
				"email":            nil,
			},
		},
		"assinante": map[string]interface{}{
			"codigo":          "123",
			"tipo":            "3",
			"servico":         "4",
			"numero_contrato": "12345678",
			"data_inicio":     "2022-01-01",
			"data_fim":        "2022-01-01",
			"numero_terminal": nil,
			"uf":              nil,
		},
		"itens": listaItens,
		"cobranca": map[string]interface{}{
			"data_competencia": "2024-06-01",
			"data_vencimento":  "2024-06-30",
			"codigo_barras":    "19872982798277298279287298728278272872872",
		},
		"informacoes_adicionais_contribuinte": "",
	}

	resp, err := nfcom.Cria(payload)

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

			respC, err := nfcom.Consulta(payload)

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

		respC, err := nfcom.Consulta(payload)

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
