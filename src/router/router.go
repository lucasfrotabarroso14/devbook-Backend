package router

import (
	"github.com/gorilla/mux"
	"api/src/router/rotas"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
	
}