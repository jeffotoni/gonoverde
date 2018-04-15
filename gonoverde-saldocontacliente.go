/*
* Go Library (C) 2018 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     Noverde
* @package     main
* @author      @jeffotoni
* @size        15/04/2018
 */

package gonoverde

import (
	"bufio"
	"fmt"
	. "github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gonoverde/gbolt"
	"log"
	"os"
	"strings"
)

var errs error

// metodo resposnsavel por gerar os saldos das contas
// dos clientes, ler suas transacoes e fazer o calculo
// de saldo e apresentar na tela
func SaldoContaCliente(ContasFile, TransFile string) {

	// Testing boltdb database
	// Start ping database
	// Creating ping ok
	gbolt.Save("Ping", "ok")

	// Testing whether it was recorded
	// and read on the boltdb, we
	// recorded a Ping and then
	// read it back.
	if gbolt.Get("Ping") != "ok" {

		log.Println("Services Error Data Base!")
		os.Exit(0)
	}

	// boltdb testado e ok
	fmt.Println(CyanCor("Bolt Testado e funcionando.."))

	fmt.Println(PurpleCor("Debita e deposita na conta do cliente"))

	// Iremos ler todo arquivo de contas.csv e coloca-lo no banco noSql boltdb
	// Apos gerar o banco de dados contas, iremos abrir e percorrer o arquivo
	// transacoes.csv, neste caso não iremos precisar gravar no banco porque
	// o arquivo encontra-se ordenado por idconta, os saldos dos clientes
	// estão ordenados ...
	// Tudo correndo bem teremos todos os saldos no banco já no formato
	// float da forma que precisamos para calcular
	errs = LerFileSaveDb(ContasFile)

	if errs != nil {

		fmt.Println("Error ao ler aquivo " + ContasFile + " não poderemos continuar!")
		log.Println(errs)
		os.Exit(0)
	}

	// agora vamos ler as transacoes e fazer os calculos para apresentar na tela
	// para ler as transacoes iremos ler linha a linha para não estourar nossa
	// memória com milhares de registros
	// Vamos quebrar em blocos as transacoes manter em um vetor em memória
	// somente o grupo de conta que estiver lendo no momento do inicio a fim
	// e depois vamos varrer este vetor da conta para efetuar os calculos
	// terminando, limpamos o vetor e inicamos tudo novamente com outra conta
	// ou seja manteremos em memória somente uma conta por vez para que possamos
	// fazer os calculos de saldo corretamente e de forma segura
	// para sabermos o saldo inicial iremos buscar no banco
	errs = CalcularSaldoTransacoes(TransFile)

	if errs != nil {

		fmt.Println("Error ao ler aquivo " + TransFile + ", isto impossibilita de fazer os calculos de saldo!")
		log.Println(errs)
		os.Exit(0)
	}
}

// Esta funcao ira ler o arquivo linha a linha
// não irá jogar tudo na memória, está otimizado
// para salvar linha a linha no banco boltdb
// o banco foi criado para otimizar a leitura
// que iremos precisar fazer quando formos
// ler as transacoes do cliente
func LerFileSaveDb(ContasFile string) error {

	// variaveis declaradas para evitar declaracoes dentro do loop
	var linha, idConta, stringSaldo, stringSaldoBody, stringSaldoDecimal string
	var Saldo float64
	var vetorConta []string

	// Abre o arquivo
	arquivo, err := os.Open(ContasFile)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return err
	}

	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	scanner := bufio.NewScanner(arquivo)

	fmt.Println(YellowCor("Lendo " + ContasFile + " e salvando no banco!"))

	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		if linha != "" {

			// vamos transformar o valor em decimal
			// sera um float para que possamos
			// fazer os calculos
			vetorConta = strings.Split(linha, ",")

			// get conta
			idConta = vetorConta[0]

			// get saldo
			stringSaldo = vetorConta[1]

			// colocar ponto nas duas ultimas posicoes
			// gerando casas decimais da string

			// somente o decimal sem as casas decimais
			stringSaldoBody = stringSaldo[:len(stringSaldo)-2]

			// somente o algarismo apos a virgula casas decimais
			stringSaldoDecimal = Substr(stringSaldo, len(stringSaldo)-2, 2)

			// gerando o saldo em float, com as casas decimais para efetuar os calculos
			Saldo = StringToSaldoWithDecimal(stringSaldoBody, stringSaldoDecimal)

			// convertendo float para string,
			// float com duas casas decimais
			SaldoFloatString := FloatToString(Saldo, 2)

			// salvar nova banco idConta => Saldo
			gbolt.Save(idConta, SaldoFloatString)
		}
	}

	// encontrando
	fmt.Println("Contas Salvas com sucesso!")

	return scanner.Err()
}

func CalcularSaldoTransacoes(TransFile string) error {

	// variaveis declaradas
	// para evitar declaracoes
	// dentro do loop
	var linha string

	// Abre o arquivo
	arquivo, err := os.Open(TransFile)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return err
	}

	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	scanner := bufio.NewScanner(arquivo)

	fmt.Println(YellowCor("Lendo " + TransFile + " e efetuando os Cálculos de Saldo!"))

	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		fmt.Println(linha)
	}

	return scanner.Err()
}
