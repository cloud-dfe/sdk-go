package main

import (
	"encoding/json"
	"fmt"
	"time"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOiJ0b2tlbl9leGVtcGxvIiwidXNyIjoidGsiLCJ0cCI6InRrIn0.Tva_viCMCeG3nkRYmi_RcJ6BtSzui60kdzIsuq5X-sQ"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false, "2")

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	nfse := sdk_cloud_dfe.Nfse(config)

	payload := map[string]any{
		"numero":       "",
		"serie":        "",
		"tipo":         "",
		"status":       "",
		"data_emissao": "",

		"tomador": map[string]any{
			"cnpj":         "",
			"cpf":          "",
			"im":           "",
			"razao_social": "",
			"endereco": map[string]any{
				"logradouro":       "",
				"numero":           "",
				"complemento":      "",
				"bairro":           "",
				"codigo_municipio": "",
				"uf":               "",
				"cep":              "",
			},
		},

		"servico": map[string]any{
			"endereco_local_prestacao": map[string]any{
				"codigo_municipio":           "",
				"codigo_municipio_prestacao": "",
				"codigo_pais":                "",
			},
			"codigo":                        "",
			"codigo_tributacao_municipio":   "",
			"discriminacao":                 "",
			"valor_servicos":                "",
			"valor_desconto_incondicionado": "",

			"tributos_municipais": map[string]any{
				"iss_retido":             "",
				"responsavel_retencao":   "",
				"valor_base_calculo_iss": "",
				"aliquota_iss":           "",
				"valor_iss":              "",
			},

			"tributos_nacionais": map[string]any{
				"valor_pis":    "",
				"valor_cofins": "",
				"valor_inss":   "",
				"valor_ir":     "",
				"valor_csll":   "",
				"valor_outras": "",
			},
		},
	}

	resp, err := nfse.Cria(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	if resp["sucesso"].(bool) {
		chave := resp["chave"].(string)
		time.Sleep(15 * time.Second)

		payload := map[string]interface{}{
			"chave": chave,
		}

		respC, err := nfse.Consulta(payload)
		if err != nil {
			fmt.Println(err)
		}

		if respC["codigo"].(int) != 5023 {
			if respC["sucesso"].(bool) {
				// autorizado
				jsonData, err := json.Marshal(respC)
				if err != nil {
					fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
				}
				fmt.Println(string(jsonData))
			} else {
				// rejeitado
				jsonData, err := json.Marshal(respC)
				if err != nil {
					fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
				}
				fmt.Println(string(jsonData))
			}
		} else {
			// nota em processamento
			// recomendado que seja utilizado o método de consulta manual ou utilizando o webhook
			jsonData, err := json.Marshal(respC)
			if err != nil {
				fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
			}
			fmt.Println(string(jsonData))
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

		respC, err := nfse.Consulta(payload)

		if err != nil {
			fmt.Println(err)
		}

		if respC["codigo"].(int) != 5023 {

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
		fmt.Printf("Erro ao converter mapa para JSON: %v \n", err)
	}

	fmt.Println(string(jsonData))

}
