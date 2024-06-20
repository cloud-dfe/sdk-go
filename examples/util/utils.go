package main

import (
	"fmt"
	"log"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	// Exemplo de Ler arquivo

	fileContent, err := sdk_cloud_dfe.ReadFile("caminho_do_arquivo.xml")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	fmt.Println("\nFile: ", fileContent)

	// Exemplo de Encode

	encodedContent := sdk_cloud_dfe.Encode(fileContent)
	fmt.Println("\nEncoded :", encodedContent)

	// Exemplo de Decode

	decodedContent, err := sdk_cloud_dfe.Decode(encodedContent)
	if err != nil {
		log.Fatalf("Erro ao decodificar o conteúdo: %v", err)
	}
	fmt.Println("\nDecoded :", decodedContent)
}
