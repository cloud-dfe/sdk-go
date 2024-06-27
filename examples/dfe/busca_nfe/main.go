package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjE0LCJ1c3IiOjgsInRwIjoyLCJpYXQiOjE2NzIyNTAzMzV9.TY8-eAg6gUFSo55epFL-UoPTD3XAUJMl8SxUcAsCr4o"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	dfe := sdk_cloud_dfe.Dfe(config)

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo.")
	}

	payload := map[string]interface{}{
		"periodo": "2020-10",
		"data":    "2020-10-15",
		"cnpj":    "06338788000127",
	}

	resp, err := dfe.BuscaNfe(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	if resp["sucesso"].(bool) {

		docs := resp["docs"].(map[string]map[string]interface{})

		for _, doc := range docs {
			chave := doc["chave"]
			fmt.Println(chave)
		}

		eventos := resp["eventos_proprios"].(map[string]map[string]interface{})

		for _, evento := range eventos {
			chave := evento["chave"]
			fmt.Println(chave)
		}
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
