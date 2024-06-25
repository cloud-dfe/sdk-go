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

	cte := sdk_cloud_dfe.Cte(config)

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo.")
	}

	payload := map[string]interface{}{
		"chave": "50000000000000000000000000000000000000000000",
		"correcoes": []map[string]interface{}{
			{
				"grupo_corrigido": "ide",
				"campo_corrigido": "natOp",
				"valor_corrigido": "PRESTACAO DE SERVIÇO",
			},
		},
	}

	resp, err := cte.Correcao(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
