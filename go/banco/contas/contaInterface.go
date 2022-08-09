package contas


type ContaInterface interface {
	Sacar(valorDoSaque float64) string
	Depositar(valorDoDeposito float64) (string, float64) 
	Transferir(valorDaTransferencia float64, contaDestino ContaInterface) bool 
	ObterSaldo() float64 
}