package main

import (
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go"
)

func main() {

	config := sdk_cloud_dfe.NewBase("Teste", sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)
	config2 := sdk_cloud_dfe.NewClient(config)
	config3, err := sdk_cloud_dfe.NewService(config2)

	if err != nil {
		println("Ocorreu um erro")
	}

	service := sdk_cloud_dfe.SetService(config3)

	exemplo := service.Request()

	fmt.Println(exemplo)
}
