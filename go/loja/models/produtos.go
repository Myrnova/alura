package models

import (
	"log"
	"loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		log.Println("Erro ao buscar os dados: ", err)
	}

	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		p := Produto{}
		selectDeTodosOsProdutos.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			log.Println("Erro ao inserir na struct os dados: ", err)
		}
		produtos = append(produtos, p)
	}

	return produtos
}

func BuscaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	defer db.Close()

	linha, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		log.Println("Erro ao buscar os dados: ", err)
	}

	defer linha.Close()

	produto := Produto{}

	if linha.Next() {
		if err := linha.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
		); err != nil {
			log.Println("Erro ao buscar produto: ", err)
		}
	}

	return produto
}

func CriarNovoProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()

	statement, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println("Erro ao preparar a query de inserção: ", err)
	}

	defer statement.Close()

	if _, err = statement.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade); err != nil {
		log.Println("Erro ao inserir os dados: ", err)
	}
	defer db.Close()

}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	statement, err := db.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		log.Println("Erro ao preparar a query de inserção: ", err)
	}

	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		log.Println("Erro ao deletar os dados: ", err)
	}
	defer db.Close()

}

func EditaProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()

	statement, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")

	if err != nil {
		log.Println("Erro ao preparar a query de inserção: ", err)
	}

	defer statement.Close()

	if _, err = statement.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id); err != nil {
		log.Println("Erro ao atualizar os dados: ", err)
	}
	defer db.Close()
}
