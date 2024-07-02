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

	xml, err := sdk_cloud_dfe.ReadFile("caminho_do_arquivo.xml")

	if err != nil {
		fmt.Println("Erro ao tentar abrir o arquivo XML")
	}

	fileXmlBase64 := sdk_cloud_dfe.Encode(xml)

	payload := map[string]interface{}{
		"chave": "50000000000000000000000000000000000000000000",
		"xml":   fileXmlBase64,
	}

	resp, err := nfse.Conflito(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
