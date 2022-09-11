package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectionString = ""
	Port                     = 0
	SecretKey                = ""
)

func LoadEnvs() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 3333
	}

	DatabaseConnectionString = fmt.Sprintf("%s:%s@tcp(containers-us-west-25.railway.app:5894)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = os.Getenv("SECRET_KEY")
}
