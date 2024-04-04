/*Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e para as funções:

A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e senha
E devem implementar as seguintes funcionalidades:

mudar o nome: me permite mudar o nome e sobrenome
mudar a idade: me permite mudar a idade
mudar e-mail: me permite mudar o e-mail
mudar a senha: me permite mudar a senha*/

package ex01

import "fmt"

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
}

func (u *Usuario) setNome(nome string) {
	u.Nome = nome
}

func (u *Usuario) setSobrenome(sobrenome string) {
	u.Sobrenome = sobrenome
}

func (u *Usuario) setIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) setEmail(email string) {
	u.Email = email
}

func (u *Usuario) setSenha(senha string) {
	u.Senha = senha
}

func (u Usuario) toString() string {
	return fmt.Sprintf("%s,%s,%d,%s,%s", u.Nome, u.Sobrenome, u.Idade, u.Email, u.Senha)
}

func ExercicioUm() {
	fmt.Println("EXERCÍCIO 1")
	usuario := Usuario{
		Nome:      "João",
		Sobrenome: "Silva",
		Idade:     30,
		Email:     "joao.silva@gmail.com",
		Senha:     "123456",
	}
	fmt.Println(usuario.toString())
	usuario.setNome("Maria")
	usuario.setSobrenome("Santos")
	usuario.setIdade(25)
	usuario.setEmail("maria.santos@gmail.com")
	usuario.setSenha("654321")
	fmt.Println(usuario.toString())
}
