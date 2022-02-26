package main

import (
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	contaDoAndre := contas.ContaPoupanca{}
	contaDoAndre.Depositar(100)
	PagarBoleto(&contaDoAndre, 60)

	fmt.Println(contaDoAndre.ObterSaldo())

	contaDoJunior := contas.ContaCorrente{}
	contaDoJunior.Depositar(500)
	PagarBoleto(&contaDoJunior, 200)

	fmt.Println(contaDoJunior.ObterSaldo())
}
