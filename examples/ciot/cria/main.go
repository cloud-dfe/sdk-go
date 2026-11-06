package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOiJ0b2tlbl9leGVtcGxvIiwidXNyIjoidGsiLCJ0cCI6InRrIn0.Tva_viCMCeG3nkRYmi_RcJ6BtSzui60kdzIsuq5X-sQ"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	ciot := sdk_cloud_dfe.Ciot(config)

	payload := map[string]interface{}{
		"identificacao": "VIAGEM-0001",
		"tipo_operacao": "1",
		"data_inicio":   "2026-11-10",
		"data_fim":      "2026-11-12",
		"valor_frete":   "4550.92",
		"contratante": map[string]interface{}{
			"cnpj":  "11222333000181",
			"nome":  "TRANSPORTADORA EXEMPLO LTDA",
			"rntrc": "12345678",
		},
		"contratado": map[string]interface{}{
			"cnpj":  "22333444000181",
			"nome":  "TRANSPORTADOR CONTRATADO LTDA",
			"rntrc": "87654321",
		},
		"destinatario": map[string]interface{}{
			"cnpj": "44555666000199",
			"nome": "DESTINATARIO EXEMPLO LTDA",
		},
		"motorista": map[string]interface{}{
			"cpf":  "12345678909",
			"nome": "MOTORISTA EXEMPLO",
		},
		"veiculos": []interface{}{
			map[string]interface{}{
				"placa": "BCD2E34",
				"eixos": "2",
				"rntrc": "87654321",
			},
		},
		"carga": map[string]interface{}{
			"tipo":     "05",
			"natureza": "8414",
			"peso":     "1520",
		},
		"viagens": []interface{}{
			map[string]interface{}{
				"quantidade": "1",
				"distancia":  "560",
				"origem": map[string]interface{}{
					"cep":              "85000000",
					"codigo_municipio": "4121604",
					"municipio":        "Renascença",
					"uf":               "PR",
				},
				"destino": map[string]interface{}{
					"cep":              "88300000",
					"codigo_municipio": "4208203",
					"municipio":        "Itajaí",
					"uf":               "SC",
				},
			},
		},
		"pagamentos": []interface{}{
			map[string]interface{}{
				"forma":      "1",
				"tipo":       "5",
				"cnpj":       "22333444000181",
				"chave_pix":  "22333444000181",
				"vencimento": "2026-11-12",
				"valor":      "4550.92",
			},
		},
	}

	resp, err := ciot.Cria(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
