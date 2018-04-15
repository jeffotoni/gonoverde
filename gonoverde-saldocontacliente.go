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
	"time"
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
	// errs = LerFileSaveDb(ContasFile)

	// if errs != nil {

	// 	fmt.Println("Error ao ler aquivo " + ContasFile + " não poderemos continuar!")
	// 	log.Println(errs)
	// 	os.Exit(0)
	// }

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
	var linha, idConta, SaldoInicialFloatString string

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

	fmt.Println(YellowCor("... Lendo " + ContasFile + " e salvando no banco! ..."))
	// RuneTime()

	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		if linha != "" {

			idConta, SaldoInicialFloatString = IdContaSaldoString(linha)

			// salvar nova banco idConta => Saldo
			gbolt.Save(idConta, SaldoInicialFloatString)
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
	var linha, idConta, ValorTransacaoStr, idContaTemp, SaldoIString string

	var SaldoInicialFloat, SaldoFloatTotal, ValorTransacaoFloat float64

	var VetorTransacao []float64

	var j int

	// set inicio
	idContaTemp = ""
	j = 0
	SaldoFloatTotal = 0

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

	// lendo arquivos e mostrando na tela
	fmt.Println(YellowCor("..... Lendo " + TransFile + " e efetuando os Cálculos de Saldo! ....."))
	fmt.Println("")

	// RuneTime()
	// varrendo o arquivo
	for scanner.Scan() {

		// get linha
		linha = scanner.Text()

		// retornando IdConta e o valor da transacao
		idConta, ValorTransacaoStr = IdContaSaldoString(linha)

		// convertendo valor da transacao em float
		ValorTransacaoFloat = StringToFloat(ValorTransacaoStr)

		// fmt.Println("Conta: ", idConta)
		// fmt.Println("transacao: ", ValorTransacaoStr)

		// entrando
		// pela primeira
		// vez
		if j == 0 {

			// coloca no vetor
			VetorTransacao = append(VetorTransacao, ValorTransacaoFloat)

			// seta e
			// nao ira entrar
			// mais nesta condicao
			j = 1

		} else {

			// se for igual armazena
			// no vetor
			if idContaTemp == idConta {

				// preenchendo o vetor com valor da transacao
				// de um cliente especifico
				// isto só é possivel pq o arquivo está ordenado
				VetorTransacao = append(VetorTransacao, ValorTransacaoFloat)

			} else {

				// buscar saldo inicial da conta
				SaldoIString = gbolt.Get(idContaTemp)

				// trasforma string em float do saldo
				SaldoInicialFloat = StringToFloat(SaldoIString)

				// Saldo
				SaldoFloatTotal = FloatCasasDecimais(SaldoFloatTotal+SaldoInicialFloat, 2)

				// saldo da conta

				fmt.Println(YellowCor("Saldo Inicial"))
				fmt.Println(SaldoInicialFloat)

				//fmt.Println(SaldoInicialFloat)
				//os.Exit(0)
				// Encontrou outro codigo cliente
				// Percorre o vetor faz os calculos do vetor e do cliente anterior
				// Limpa o vetor e inicia seu preenchimento novamente
				// com um novo codigo cliente
				// Calculo do Saldo
				for _, Tvalor := range VetorTransacao {

					// O saldo de uma conta deve ser calculado a partir de seu saldo inicial, aplicando cada uma das
					// transações relacionadas a esta conta. Depósitos devem aumentar o saldo da conta e débitos
					// devem reduzir esse mesmo saldo, na medida do valor da transação.

					// Uma conta pode​ assumir um valor negativo e não existe limite inferior para o saldo da conta.
					// Contudo, cada transação de débito que termina deixando o saldo da conta negativo implica
					// uma multa de R$ 5,00​ a ser descontada imediatamente. Esta multa se aplica independente da
					// conta se encontrar ou não com saldo negativo antes da transação, mas não se aplica se a
					// transação for um depósito

					// transacao negativa
					if Tvalor < 0 {

						fmt.Println(Tvalor)

					} else {

						fmt.Println(Tvalor)

					}

					// converte para duas casas decimais
					SaldoFloatTotal = FloatCasasDecimais(SaldoFloatTotal+(Tvalor), 2)
				}

				// total Saldo
				fmt.Println(YellowCor("saldo total"))
				fmt.Println(SaldoFloatTotal)
				fmt.Println("")

				time.Sleep(time.Second * 10)

				//fmt.Println("Conta fim: ", idConta)
				//os.Exit(0)

				// limpar vetor
				// iniciar novamente
				VetorTransacao = []float64{}
				SaldoFloatTotal = 0

				// carregando o vetor com idConta do proximo cliente
				VetorTransacao = append(VetorTransacao, ValorTransacaoFloat)
			}

		}

		// pegar o idConta e o valor da Transacao
		//fmt.Println(linha)
		//fmt.Println(idConta + "," + ValorTransacaoStr)
		//fmt.Println("")

		idContaTemp = idConta

		//fmt.Println(gbolt.Get())_
	}

	return scanner.Err()
}
