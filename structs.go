package sdk_cloud_dfe

import "fmt"

type Ambiente int
type BaseUri string

const (
	AMBIENTE_PRODUCAO    Ambiente = 1
	AMBIENTE_HOMOLOGACAO Ambiente = 2
	API_PRODUCAO         BaseUri  = "https://api.integranotas.com.br/v1"
	API_HOMOLOGACAO      BaseUri  = "https://hom-api.integranotas.com.br/v1"
)

type configBase struct {
	Token    string
	Ambiente Ambiente
	Timeout  int
	Port     int
	Debug    bool
}

func NewBase(token string, ambiente Ambiente, timeout int, port int, debug bool) configBase {

	return configBase{
		Token:    token,
		Ambiente: ambiente,
		Timeout:  timeout,
		Port:     port,
		Debug:    debug,
	}
}

type configClient struct {
	Token    string
	Ambiente Ambiente
	Timeout  int
	Port     int
	Debug    bool
}

func NewClient(config configBase) configClient {

	token := config.Token
	ambiente := config.Ambiente
	timeout := config.Timeout
	port := config.Port
	debug := config.Debug

	return configClient{
		Token:    token,
		Ambiente: ambiente,
		Timeout:  timeout,
		Port:     port,
		Debug:    debug,
	}
}

type configRequest struct {
	BaseUri BaseUri
	Token   string
	Timeout int
	Port    int
	Debug   bool
}

func NewRequest(config configClient) (configRequest, error) {
	token := config.Token
	ambiente := config.Ambiente
	timeout := config.Timeout
	port := config.Port
	debug := config.Debug

	if ambiente != AMBIENTE_PRODUCAO && ambiente != AMBIENTE_HOMOLOGACAO {
		return configRequest{}, nil

	} else {
		if ambiente == AMBIENTE_PRODUCAO {
			baseUri := API_PRODUCAO

			return configRequest{
				BaseUri: baseUri,
				Token:   token,
				Timeout: timeout,
				Port:    port,
				Debug:   debug,
			}, nil

		} else {
			baseUri := API_HOMOLOGACAO

			return configRequest{
				BaseUri: baseUri,
				Token:   token,
				Timeout: timeout,
				Port:    port,
				Debug:   debug,
			}, nil
		}
	}
}

func PrintConfig(config configBase) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Ambiente: %d\n", config.Ambiente)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}

func PrintConfig2(config configClient) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Ambiente: %d\n", config.Ambiente)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}

func PrintConfig3(config configRequest) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Base URI: %s\n", config.BaseUri)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}
