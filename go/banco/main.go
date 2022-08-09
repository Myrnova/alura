package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)
func PagarBoleto(conta contas.ContaInterface, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}


func main() {
	contaDoGuilherme := contas.ContaCorrente{
		Titular: clientes.Titular{ Nome: "Guilherme"},
		NumeroAgencia: 589,
		NumeroConta: 123456,
	}
	contaDaBruna := contas.ContaCorrente{
		Titular: clientes.Titular{"Bruna","123.111.121.31",	"Desenvolvedor"},
		NumeroAgencia: 222,
		NumeroConta:   111222,
	}

	var contaDaCris *contas.ContaCorrente = new(contas.ContaCorrente)
	contaDaCris.Titular = clientes.Titular{Nome: "Cris"}
	contaDaCris.Depositar(500)

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(*contaDaCris)

	status, valor := contaDoGuilherme.Depositar(100)
	fmt.Println(status, valor)
	fmt.Println(contaDoGuilherme.Sacar(10))

	statusTransferencia := contaDaBruna.Transferir(100, &contaDoGuilherme) //& passa o endereço de memoria

	fmt.Println(statusTransferencia)

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	PagarBoleto(&contaDoGuilherme, 60)
	fmt.Println(contaDoGuilherme.ObterSaldo())

}
