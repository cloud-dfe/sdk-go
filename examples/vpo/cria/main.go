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
		"numero": "1001",
		"identificacao": "VIAGEM-77",
		"emissor": "SEMPARAR",
		"data_inicio": "2026-11-10 08:00:00",
		"data_fim": "2026-11-12 18:00:00",
		"rota": map[string]interface{}{
			"id": "58222",
			"indice": "0",
			"distancia_km": "45.279",
			"tipo_combustivel": "DIESEL",
		},
		"veiculo": map[string]interface{}{
			"placa": "BCD2E34",
			"categoria": "7",
		},
		"transportador": map[string]interface{}{
			"cnpj": "22333444000181",
			"rntrc": "87654321",
			"nome": "TRANSPORTADOR EXEMPLO LTDA",
		},
		"cte": "1234",
	}

	resp, err := vpo.Cria(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
