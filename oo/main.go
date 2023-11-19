package main

import "fmt"

// Definindo uma struct chamada 'Pessoa'
type Pessoa struct {
	Nome  string
	Idade int
}

// Método associado à struct Pessoa
func (p Pessoa) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

// Método associado à struct Pessoa que modifica o valor da idade
func (p *Pessoa) Aniversario() {
	p.Idade++
}

func main() {
	// Criando uma instância da struct Pessoa
	pessoa := Pessoa{
		Nome:  "Alice",
		Idade: 25,
	}

	// Chamando o método Apresentar
	pessoa.Apresentar()

	// Chamando o método Aniversario
	pessoa.Aniversario()

	// Chamando novamente o método Apresentar após o aniversário
	pessoa.Apresentar()
}