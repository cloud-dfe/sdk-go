package main

import (
	"encoding/json"
	"fmt"
	"time"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOiJ0b2tlbl9leGVtcGxvIiwidXNyIjoidGsiLCJ0cCI6InRrIn0.Tva_viCMCeG3nkRYmi_RcJ6BtSzui60kdzIsuq5X-sQ"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	gnre := sdk_cloud_dfe.Gnre(config)

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo.")
	}

	payload := map[string]interface{}{
		"numero":                    "6",
		"uf_favoverida":             "RO",
		"ie_emitente_uf_favorecida": nil,
		"tipo":                      "0",
		"valor":                     12.55,
		"data_pagamento":            "2022-05-22",
		"identificador_guia":        "12345",
		"receitas": []map[string]interface{}{
			{
				"codigo":               "100102",
				"detalhamento":         nil,
				"data_vencimento":      "2022-05-22",
				"convenio":             "Convênio ICMS 142/18",
				"numero_controle":      "1",
				"numero_controle_fecp": nil,
				"documento_origem": map[string]interface{}{
					"numero": "000000001",
					"tipo":   "10",
				},
				"produto": nil,
				"referencia": map[string]interface{}{
					"periodo": "0",
					"mes":     "05",
					"ano":     "2022",
					"parcela": nil,
				},
				"valores": []map[string]interface{}{
					{
						"valor": 12.55,
						"tipo":  "11",
					},
				},
				"contribuinte_destinatario": map[string]interface{}{
					"cnpj":  nil,
					"cpf":   nil,
					"ie":    nil,
					"razao": nil,
					"ibge":  nil,
				},
				"extras": []map[string]interface{}{
					{
						"codigo":   "52",
						"conteudo": "32220526434850000191550100000000011015892724",
					},
				},
			},
		},
	}

	resp, err := gnre.Cria(payload)

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

			respC, err := gnre.Consulta(payload)

			if err != nil {
				print(err)
			}

			if respC["codigo"].(int) == 5023 {

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

	} else if resp["codigo"].(float64) == 5008 {
		chave := resp["chave"].(string)

		payload := map[string]interface{}{
			"chave": chave,
		}

		respC, err := gnre.Consulta(payload)

		if err != nil {
			fmt.Println(err)
		}

		if respC["codigo"].(float64) == 5023 {
			if respC["sucesso"].(bool) {
				jsonData, err := json.Marshal(respC)
				if err != nil {
					fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
				}
				fmt.Println(string(jsonData))

			} else {
				jsonData, err := json.Marshal(respC)
				if err != nil {
					fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
				}
				fmt.Println(string(jsonData))
			}
		} else {
			jsonData, err := json.Marshal(respC)
			if err != nil {
				fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
			}
			fmt.Println(string(jsonData))
		}

	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
