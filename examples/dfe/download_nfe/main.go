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
		"chave": "50000000000000000000000000000000000000000000",
	}

	resp, err := dfe.DownloadNfe(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	if resp["sucesso"].(bool) {
		doc := resp["doc"].(map[string]interface{})

		xmlBase64 := doc["xml"].(string)
		pdfBase64 := doc["pdf"].(string)

		xml, err := sdk_cloud_dfe.Decode(xmlBase64)

		if err != nil {
			fmt.Println("Erro ao tentar decodificar")
		}

		pdf, err := sdk_cloud_dfe.Decode(pdfBase64)

		if err != nil {
			fmt.Println("Erro ao tentar decodificar")
		}

		fmt.Println(xml)

		fmt.Println(pdf)

	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
