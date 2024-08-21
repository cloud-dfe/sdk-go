package main

import (
	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOiJ0b2tlbl9leGVtcGxvIiwidXNyIjoidGsiLCJ0cCI6InRrIn0.Tva_viCMCeG3nkRYmi_RcJ6BtSzui60kdzIsuq5X-sQ"
	payload := map[string]interface{}{
		"signature": "ojAm16Ye1cnnSxIM1D/8uUZROFYMitC6YleumaQx5W5IstqC1pdjvlact1q6LbE9f0OEjbtXZeVPYK/PtOfTmw==",
	}

	isValid, err := sdk_cloud_dfe.IsValid(token, payload)

	if err != nil {
		println("Erro:", err.Error())
	} else if isValid {
		println("Assinatura válida")
	} else {
		println("Assinatura inválida")
	}
}
