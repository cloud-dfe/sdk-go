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

	nfcom := sdk_cloud_dfe.Nfcom(config)

	payload := map[string]interface{}{
		"numero_inicial": 1210,
		"numero_final":   1210,
		"serie":          1,
		"data_inicial":   "2019-12-01",
		"data_final":     "2019-12-31",
		"cancel_inicial": "2019-12-01",
		"cancel_final":   "2019-12-31",
		"status":         "1",
	}

	resp, err := nfcom.Busca(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
