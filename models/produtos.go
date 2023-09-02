package models

import (
	"Web/db"
)

type Produto struct {
	Id         int     `gorm:"primaryKey"`
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaComBD()

	var produtos []Produto
	result := db.Order("id asc").Find(&produtos)
	if result.Error != nil {
		panic(result.Error)
	}

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBD()

	produto := Produto{
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}

	result := db.Create(&produto)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeletaProduto(id string) {
	db := db.ConectaComBD()

	result := db.Delete(&Produto{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBD()

	var produto Produto
	result := db.First(&produto, id)
	if result.Error != nil {
		panic(result.Error)
	}

	return produto
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBD()

	produto := Produto{
		Id:         id,
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}

	result := db.Save(&produto)
	if result.Error != nil {
		panic(result.Error)
	}
}
