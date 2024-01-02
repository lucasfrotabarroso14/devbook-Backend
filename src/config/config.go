package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	//string conexao com o mysql
	StringConexaoBanco = " "
	//porta  onde a api vai estar rodando
	Porta = 0

	SecretKey []byte
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {

	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		//VAO SUBSTITUIR A STRING DE CONEXAO ONDE TEM O s
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
