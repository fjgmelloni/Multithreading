package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fjgmelloni/fullcycle/multithreading/metrics"
	"net/http"
	"strings"
	"time"
)

type CepResponse struct {
	Source     string `json:"source"`
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func CepHandler(w http.ResponseWriter, r *http.Request) {
	cep := strings.TrimPrefix(r.URL.Path, "/cep/")
	if cep == "" {
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}

	fmt.Printf("[%s] CEP requisitado: %s\n", time.Now().Format("15:04:05.000"), cep)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result := make(chan CepResponse, 2)

	go func() {
		fmt.Printf("[%s] Iniciando chamada para BrasilAPI\n", time.Now().Format("15:04:05.000"))
		if res := fetchFromBrasilAPI(ctx, cep); res != nil {
			select {
			case result <- *res:
			case <-ctx.Done():
			}
		}
	}()

	go func() {
		fmt.Printf("[%s] Iniciando chamada para ViaCEP\n", time.Now().Format("15:04:05.000"))
		if res := fetchFromViaCEP(ctx, cep); res != nil {
			select {
			case result <- *res:
			case <-ctx.Done():
			}
		}
	}()

	select {
	case r := <-result:
		fmt.Printf("[%s] Resposta recebida da %s\n", time.Now().Format("15:04:05.000"), r.Source)
		if r.Source == "BrasilAPI" {
			metrics.IncrementBrasilAPI()
		} else {
			metrics.IncrementViaCEP()
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(r)
	case <-ctx.Done():
		fmt.Printf("[%s] Timeout ao consultar o CEP %s\n", time.Now().Format("15:04:05.000"), cep)
		http.Error(w, "Timeout: nenhuma API respondeu a tempo", http.StatusGatewayTimeout)
	}
}

func fetchFromBrasilAPI(ctx context.Context, cep string) *CepResponse {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return &CepResponse{
		Source:     "BrasilAPI",
		Cep:        data["cep"].(string),
		Logradouro: data["street"].(string),
		Bairro:     data["neighborhood"].(string),
		Localidade: data["city"].(string),
		Uf:         data["state"].(string),
	}
}

func fetchFromViaCEP(ctx context.Context, cep string) *CepResponse {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return &CepResponse{
		Source:     "ViaCEP",
		Cep:        data["cep"].(string),
		Logradouro: data["logradouro"].(string),
		Bairro:     data["bairro"].(string),
		Localidade: data["localidade"].(string),
		Uf:         data["uf"].(string),
	}
}
