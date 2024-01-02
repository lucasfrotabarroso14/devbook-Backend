package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em json para a requsicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}

	}

}

func Erro(w http.ResponseWriter, statusCide int, erro error) {
	JSON(w, statusCide, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})

}
