package database

import {
	  "gorm.io/driver/postgres"
	  "gorm.io/gorm"
}

var {
	DB *gorm.db
	err error
}

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(stringDeConexao)
	if err != nil {
		log.Panic("Erro ao conectar com Banco de Dados")
	}
}
