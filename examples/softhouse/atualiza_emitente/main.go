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

	softhouse := sdk_cloud_dfe.Softhouse(config)

	payload := map[string]interface{}{
		"nome":                     "EMPRESA TESTE",
		"razao":                    "EMPRESA TESTE",
		"cnae":                     "12369875",
		"crt":                      "1", // Regime tributário
		"ie":                       "12369875",
		"im":                       "12369875",
		"suframa":                  "12369875",
		"csc":                      "...", // token para emissão de NFCe
		"cscid":                    "000001",
		"tar":                      "C92920029-12", // tar BPe
		"login_prefeitura":         nil,
		"senha_prefeitura":         nil,
		"client_id_prefeitura":     nil,
		"client_secret_prefeitura": nil,
		"telefone":                 "46998895532",
		"email":                    "empresa@teste.com",
		"rua":                      "TESTE",
		"numero":                   "1",
		"complemento":              "NENHUM",
		"bairro":                   "TESTE",
		"municipio":                "CIDADE TESTE", // IBGE
		"cmun":                     "5300108",      // IBGE
		"uf":                       "PR",           // IBGE
		"cep":                      "85000100",
		"logo":                     "useyn56j4mx35m5j6_JSHh734khjd...saasjda", // BASE 64
		"webhook":                  "https://seusite.com.br/notifications",
	}

	resp, err := softhouse.AtualizaEmitente(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
