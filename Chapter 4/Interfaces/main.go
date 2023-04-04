package main

import ("fmt")

type Humano interface {
	Fala(texto string)
	Sente(sentimento string)
}

type Robo struct {
	Nome string
	DataFabricao string
}

func (r *Robo) Fala(texto string) {
	fmt.Println(texto)
}

type Pessoa struct {
	Nome string
	Idade int
}

func (p *Pessoa) Fala(texto string) {
	fmt.Println(texto)
}

func (p *Pessoa) Sente(sentimento string) {
	fmt.Println("Estou sentindo", sentimento)
}

func CadastrarCPF(h Humano) {
	// h.CPF = "12345678910"
	fmt.Println("CPF Cadastrado")
}

func main() {

	p := &Pessoa{}
	r := &Robo{}

	p.Fala("Olá eu sou uma pessoa")
	r.Fala("Olá eu sou um robo")

	CadastrarCPF(p)
	CadastrarCPF(r)
}
