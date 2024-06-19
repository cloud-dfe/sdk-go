package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjE0LCJ1c3IiOjgsInRwIjoyLCJpYXQiOjE2NzIyNTAzMzV9.TY8-eAg6gUFSo55epFL-UoPTD3XAUJMl8SxUcAsCr4o"

	config := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)
	config2 := sdk_cloud_dfe.NewClient(config)
	service, err := sdk_cloud_dfe.NewService(config2)

	if err != nil {
		println("Ocorreu um erro")
	}

	exemplo, err := service.Request(http.MethodGet, "/emitente", nil)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	jsonData2, err := json.Marshal(exemplo)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData2))
}
