/*Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o mesmo endereço de memória no main do programa e nas funções.


Estruturas necessárias:

Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
Produto: Nome, preço, quantidade.
Algumas funções necessárias:

Novo produto: recebe nome e preço, e retorna um produto.
Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
Deletar produtos: recebe um usuário, apaga  os produtos do usuário.*/

package ex02

import "fmt"

type Product struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type User struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []Product
}

func criaProduct(nome string, preco float64, quantidade int) *Product {
	return &Product{
		Nome:       nome,
		Preco:      preco,
		Quantidade: quantidade,
	}
}

func addProduct(u *User, p *Product) {
	u.Produtos = append(u.Produtos, *p)
}

func deleteUserProducts(u *User) {
	u.Produtos = []Product{}
}

func (p Product) toString() string {
	return fmt.Sprintf("- %10s | R$ %10.2f\n", p.Nome, p.Preco)
}

func (u User) toString() string {
	text := fmt.Sprintf("Nome: %s %s\nE-mail: %s\n", u.Nome, u.Sobrenome, u.Email)
	if len(u.Produtos) > 0 {
		text += "Produtos: \n"
		for _, produto := range u.Produtos {
			text += produto.toString()
		}
	} else {
		text += "Produtos: NENHUM"
	}
	return text
}

func ExercicioDois() {
	fmt.Println("\nEXERCÍCIO 2")
	user := User{
		Nome:      "John",
		Sobrenome: "Doe",
		Email:     "john.doe@example.com",
	}

	prod1 := criaProduct("Maça", 10.0, 5)
	prod2 := criaProduct("Banana", 20.0, 10)
	prod3 := criaProduct("Queijo", 2000.0, 10)

	addProduct(&user, prod1)
	addProduct(&user, prod2)
	addProduct(&user, prod3)
	fmt.Println(user.toString())

	deleteUserProducts(&user)
	fmt.Println(user.toString())
}
