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

	vpo := sdk_cloud_dfe.Vpo(config)

	payload := map[string]interface{}{
		"enderecos": []interface{}{
			map[string]interface{}{
				"codigo_municipio": "3543402",
				"nome_municipio": "Ribeirão Preto",
				"uf": "SP",
			},
			map[string]interface{}{
				"codigo_municipio": "3505906",
				"nome_municipio": "Batatais",
				"uf": "SP",
			},
		},
		"categoria": "7",
		"data_inicio": "2026-11-10 08:00:00",
		"tipo_combustivel": "DIESEL",
	}

	resp, err := vpo.Rota(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
