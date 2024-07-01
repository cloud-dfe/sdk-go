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
		"numero_rps_inicial": 15,
		"numero_rps_final":   15,
		"serie_rps":          "0",
		// "numero_nfse_inicial": 1210,
		// "numero_nfse_final": 1210,
		// "data_inicial": "2019-12-01",  // Autorização
		// "data_final": "2019-12-31",
		// "cancel_inicial": "2019-12-01",  // Cancelamento
		// "cancel_final": "2019-12-31"
	}

	resp, err := nfse.Busca(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
