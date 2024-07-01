package main

import (
	"fmt"
	"log"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {
	fileContent, err := sdk_cloud_dfe.ReadFile("caminho_do_arquivo.xml")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	fmt.Println("File: ", fileContent)
}
