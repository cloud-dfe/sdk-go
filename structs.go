package main

import "fmt"

type Ambiente int
type BaseUri string

const (
	AMBIENTE_PRODUCAO    Ambiente = 1
	AMBIENTE_HOMOLOGACAO Ambiente = 2
	API_PRODUCAO         BaseUri  = "https://api.integranotas.com.br/v1"
	API_HOMOLOGACAO      BaseUri  = "https://hom-api.integranotas.com.br/v1"
)

type ConfigBase struct {
	Token    string
	Ambiente Ambiente
	Timeout  int
	Port     int
	Debug    bool
}

func NewBase(token string, ambiente Ambiente, timeout int, port int, debug bool) ConfigBase {

	return ConfigBase{
		Token:    token,
		Ambiente: ambiente,
		Timeout:  timeout,
		Port:     port,
		Debug:    debug,
	}
}

type ConfigClient struct {
	Token    string
	Ambiente Ambiente
	Timeout  int
	Port     int
	Debug    bool
}

func NewClient(config ConfigBase) ConfigClient {

	token := config.Token
	ambiente := config.Ambiente
	timeout := config.Timeout
	port := config.Port
	debug := config.Debug

	return ConfigClient{
		Token:    token,
		Ambiente: ambiente,
		Timeout:  timeout,
		Port:     port,
		Debug:    debug,
	}
}

type ConfigRequest struct {
	BaseUri BaseUri
	Token   string
	Timeout int
	Port    int
	Debug   bool
}

func NewRequest(config ConfigClient) (ConfigRequest, error) {
	token := config.Token
	ambiente := config.Ambiente
	timeout := config.Timeout
	port := config.Port
	debug := config.Debug

	if ambiente != AMBIENTE_PRODUCAO && ambiente != AMBIENTE_HOMOLOGACAO {
		return ConfigRequest{}, nil

	} else {
		if ambiente == AMBIENTE_PRODUCAO {
			baseUri := API_PRODUCAO

			return ConfigRequest{
				BaseUri: baseUri,
				Token:   token,
				Timeout: timeout,
				Port:    port,
				Debug:   debug,
			}, nil

		} else {
			baseUri := API_HOMOLOGACAO

			return ConfigRequest{
				BaseUri: baseUri,
				Token:   token,
				Timeout: timeout,
				Port:    port,
				Debug:   debug,
			}, nil
		}
	}
}

func main() {
	config := NewBase("Teste", AMBIENTE_HOMOLOGACAO, 60, 443, false)
	config2 := NewClient(config)
	config3, err := NewRequest(config2)

	if err != nil {
		println("Ocorreu um erro")
	}

	printConfig(config)
	fmt.Println("")
	printConfig2(config2)
	fmt.Println("")
	printConfig3(config3)
}

func printConfig(config ConfigBase) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Ambiente: %d\n", config.Ambiente)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}

func printConfig2(config ConfigClient) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Ambiente: %d\n", config.Ambiente)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}

func printConfig3(config ConfigRequest) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Base URI: %s\n", config.BaseUri)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}
