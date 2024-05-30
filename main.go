package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	viaCepUrlPattern    = "https://viacep.com.br/ws/%s/json"
	brasilApiUrlPattern = "https://brasilapi.com.br/api/cep/v1/%s"
	defaultCep          = "04707-900"
	executeTimeout      = 1
	executeTimeoutUnit  = time.Second
)

func main() {
	cep := getCep()

	viaCepChannel := make(chan string)
	brasilApiChannel := make(chan string)

	go fetchCepWithBrasilApi(cep, brasilApiChannel)
	go fetchCepWithViaCep(cep, viaCepChannel)

	select {
	case response := <-brasilApiChannel:
		log.Printf("BrasilAPI returned first: %v\n", response)
	case response := <-viaCepChannel:
		log.Printf("ViaCEP returned first: %v\n", response)
	case <-time.After(executeTimeout * executeTimeoutUnit):
		log.Println("Execution time expired")
	}
}

func getCep() string {
	if len(os.Args) < 2 {
		return defaultCep
	}
	return os.Args[1]
}

func fetchCepWithViaCep(cep string, channel chan string) {
	brasilApiUrl := fmt.Sprintf(viaCepUrlPattern, cep)

	body, err := fetch(brasilApiUrl)
	if err != nil {
		log.Printf("There was an error when fetching ViaCEP. %s\n", err.Error())
		return
	}

	channel <- string(body)
}

func fetchCepWithBrasilApi(cep string, channel chan string) {
	brasilApiUrl := fmt.Sprintf(brasilApiUrlPattern, cep)

	body, err := fetch(brasilApiUrl)
	if err != nil {
		log.Printf("There was an error when fetching BrasilAPI. %s\n", err.Error())
		return
	}

	channel <- string(body)
}

func fetch(url string) (response []byte, err error) {
	request, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
