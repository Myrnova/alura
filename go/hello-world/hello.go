package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()
		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa")
		// }else {
		// 	fmt.Println("Não conheço este comando")
		// }

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)

		}
	}
}

func exibeIntroducao() {
	var nome = "Débora"
	//var idade int = 25
	versao := 1.1

	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //é capaz de inferir o tipo que vai vir porque a variavel que esta sendo passada é do tipo int
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	//fmt.Scanf("%d", comando)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br",
	// }

	for i := 0; i < monitoramentos; i++ {
		for indice, site := range sites {
			fmt.Println("Testando site", indice, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ao tentar acessar o site", site, err)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
		return
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code", response.StatusCode)
		registraLog(site, false)
		return
	}

}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return []string{}
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Ocorreu um erro: ", err)
			return []string{}
		}
	}

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")
}

func imprimeLog() {
	fmt.Println("Exibindo Logs...")
	fmt.Println("")
	arquivo, err := ioutil.ReadFile("log.txt") //ioutil já fecha o arquivo então não é necessário um arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))

}
