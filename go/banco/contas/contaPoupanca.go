package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

/*
Também é possível passar um valor removendo a assinatura do ponteiro (c *ContaPoupanca) para (c ContaPoupanca).
Nesse caso, uma cópia do valor de ContaPoupanca é passada para a função, sem alterar o valor do ponteiro.
Portanto, precisamos ficar atento, já que qualquer alteração que você faça em c se passar por valor não será refletida na fonte c.
*/

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo
	if valorDoSaque > 0 {
		if podeSacar {
			c.saldo -= valorDoSaque
			return "Saque realizado com sucesso"
		}
		return "Saldo insuficiente"
	}
	return "Não é possivel sacar valores negativos"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "Valor do deposito menor que zero", c.saldo
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino ContaInterface) bool { //* fala que vai receber um endereço de memória
	podeTransferir := valorDaTransferencia <= c.saldo
	if valorDaTransferencia > 0 && podeTransferir {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
