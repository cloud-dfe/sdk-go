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

	nfse := sdk_cloud_dfe.Nfse(config)

	payload := map[string]interface{}{
		"data_emissao_inicial":     "2020-01-01",
		"data_emissao_final":       "2020-01-31",
		"data_competencia_inicial": "2020-01-01",
		"data_competencia_final":   "2020-01-31",
		"tomador_cnpj":             nil,
		"tomador_cpf":              nil,
		"tomador_im":               nil,
		"nfse_numero":              nil,
		"nfse_numero_inicial":      nil,
		"nfse_numero_final":        nil,
		"rps_numero":               "15",
		"rps_serie":                "0",
		"rps_tipo":                 "1",
		"pagina":                   "1",
	}

	resp, err := nfse.Localiza(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
