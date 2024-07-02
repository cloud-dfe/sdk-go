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

	mdfe := sdk_cloud_dfe.Mdfe(config)

	payload := map[string]interface{}{
		"chave":                            "50000000000000000000000000000000000000000000",
		"codigo_municipio_carregamento":    "2408003",
		"nome_municipio_carregamento":      "Mossoró",
		"codigo_municipio_descarregamento": "5200050",
		"nome_municipio_descarregamento":   "Abadia de Goiás",
		"chave_nfe":                        "50000000000000000000000000000000000000000001",
	}

	resp, err := mdfe.Nfe(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
