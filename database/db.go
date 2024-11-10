package database

import (
	  "gorm.io/driver/postgres"
	  "gorm.io/gorm"
  )

var (
	DB *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	strDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com Banco de Dados")
	}
}
